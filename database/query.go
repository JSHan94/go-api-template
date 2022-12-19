package database

import (
	"encoding/json"
	"strings"

	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
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
	return hits, err
}

func search(client Client, indexName string, query string) (*opensearchapi.Response, error) {
	content := strings.NewReader(query)

	searchResponse, err := client.OSClient.Search(
		client.OSClient.Search.WithIndex(indexName),
		client.OSClient.Search.WithBody(content),
	)

	if err != nil {
		return nil, err
	}

	return searchResponse, err
}