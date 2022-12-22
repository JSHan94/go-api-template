package config

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/opensearch-project/opensearch-go/v2"
	"github.com/opensearch-project/opensearch-go/v2/opensearchtransport"
	requestsigner "github.com/opensearch-project/opensearch-go/v2/signer/awsv2"
)

type DatabaseConfig struct {
	OSConfig *opensearch.Config
}

// database - all DB variables
func database() *DatabaseConfig {
	config := &DatabaseConfig{}

	env()

	// if OPENSEARCH_ENDPOINTS is unset, opensearch client will try with OPENSEARCH_ENDPOINTS
	// the name of 'OPENSEARCH_ENDPOINTS' is for a sense of unity with other environmental variables.
	addrs := strings.Split(os.Getenv("OPENSEARCH_ENDPOINTS"), ",")

	useAWS := os.Getenv("USE_OPENSEARCH_AWS")

	if strings.EqualFold(useAWS, "true") {
		ctx := context.Background()
		cfg, err := awsConfig.LoadDefaultConfig(ctx)
		if err != nil {
			fmt.Printf("failed to get config from aws: %+v\n", err)
			os.Exit(10)
		}

		// Create an AWS request Signer and load AWS configuration using default config folder or env vars.
		// See https://docs.aws.amazon.com/opensearch-service/latest/developerguide/request-signing.html#request-signing-go
		signer, err := requestsigner.NewSigner(cfg)
		if err != nil {
			fmt.Printf("failed to get signer from aws: %+v\n", err)
			os.Exit(11)
		}

		config.OSConfig = &opensearch.Config{
			Addresses: addrs,
			Signer:    signer,
		}
	} else {
		config.OSConfig = &opensearch.Config{
			Addresses: addrs,
			Username:  os.Getenv("OPENSEARCH_USERNAME"),
			Password:  os.Getenv("OPENSEARCH_PASSWORD"),
			//Header:                map[string][]string{},
			//Signer:                nil,
			//CACert:                []byte{},
			//RetryOnStatus:         []int{},
			DisableRetry:         false, // default:false
			EnableRetryOnTimeout: false, // default:false

			MaxRetries: 3, // default:3
			//CompressRequestBody:   false,
			// DiscoverNodesOnStart: false, // default:false
			//DiscoverNodesInterval: 0,
			EnableMetrics: true, // default:false
			//EnableDebugLogger:     false,
			//UseResponseCheckOnly:  false,
			//RetryBackoff: func(attempt int) time.Duration {},
			//Transport: nil,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				//MaxIdleConnsPerHost: 1,
			},
			Logger: &opensearchtransport.ColorLogger{Output: os.Stderr},
			//Selector:  nil,
			//ConnectionPoolFunc: func([]*opensearchtransport.Connection, opensearchtransport.Selector) opensearchtransport.ConnectionPool {},
		}
	}

	return config
}
