package database

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	"github.com/sirupsen/logrus"
)

// Search for documents
func (client Client) Search(indexName string, query string) ([]interface{}, error) {
	res, err := search(client, indexName, query)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}
	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})
	if len(hits) == 0 {
		return nil, errors.New("no hits for the request")
	}
	return hits, err
}

func search(client Client, indexName string, query string) (*opensearchapi.Response, error) {
	content := strings.NewReader(query)

	searchResponse, err := client.OSClient.Search(
		client.OSClient.Search.WithIndex(indexName),
		client.OSClient.Search.WithBody(content),
	)
	if err != nil {
		logrus.Error("no hits for the request")
		return nil, err
	}

	return searchResponse, err
}
