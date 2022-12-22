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
