package database

import (
	gconfig "github.com/initia-labs/initia-apis/config"
	"github.com/opensearch-project/opensearch-go/v2"
	"github.com/sirupsen/logrus"
)

type Client struct {
	OSClient *opensearch.Client
}

var client *Client

// names of indices
const (
	IDX_BLOCK_BASIC   = "initia-block-idx-1"
	IDX_TX_BASIC      = "initia-tx-idx-1"
	IDX_EVENT_BASIC   = "initia-event-idx-1"
	IDX_NETWORK_BASIC = "initia-network-idx-1"
	IDX_GENERAL_BASIC = "initia-general-idx-1"
)

func NewClient(config *gconfig.Configuration) *Client {
	osclient, err := opensearch.NewClient(*config.Database.OSConfig)
	if err != nil {
		logrus.Error("failed to create opensearch client: ", err)
		return nil
	}
	return &Client{OSClient: osclient}
}

func GetClient() *Client {
	if client == nil {
		logrus.Info("db client is nil... create new client")
		config := gconfig.GetConfig()
		client = NewClient(config)
	}
	return client
}
