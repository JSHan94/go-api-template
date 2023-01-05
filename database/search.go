package database

import (
	"encoding/json"
	"strings"

	glib "github.com/initia-labs/initia-apis/lib"
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
	_, ok := r["hits"].(map[string]interface{})
	if !ok {
		return nil, glib.ErrInvalidRequest
	}
	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})
	if len(hits) == 0 {
		return nil, glib.ErrNotFound
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
		return nil, glib.ErrNotFound
	}

	return searchResponse, err
}
