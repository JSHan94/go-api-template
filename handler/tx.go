package handler

import (
	gdatabase "github.com/initia-labs/initia-apis/database"
	gmodel "github.com/initia-labs/initia-apis/database/model"
	"github.com/sirupsen/logrus"
)

func GetTxByHash(indexName string, params GetTxQueryParameter) (*gmodel.CollectedTx, error) {
	// client := gdatabase.GetClient()
	return nil, nil
}

func GetTxsByAccount(indexName string, params GetTxsQueryParameter) ([]gmodel.CollectedTxs, error) {
	if params.Limit == "" {
		logrus.Info("limit is empty, set default limit to 10")
		params.Limit = "10"
	}
	query := GetTxsByAccountQuery(params)

	return GetTxs(indexName, query)
}

func GetTxsByHeight(indexName string, params GetTxsQueryParameter) ([]gmodel.CollectedTxs, error) {
	query := GetTxsByHeightQuery(params)

	return GetTxs(indexName, query)
}

func GetTxsByOffset(indexName string, params GetTxsQueryParameter) ([]gmodel.CollectedTxs, error) {
	query := GetTxsByOffsetQuery(params)

	return GetTxs(indexName, query)
}

func GetTxs(indexName string, query string) ([]gmodel.CollectedTxs, error) {
	client := gdatabase.GetClient()

	hits, err := client.Search(indexName, query)
	if err != nil {
		return nil, err
	}

	var collectedTxs []gmodel.CollectedTxs
	err = DecodeMany(hits, &collectedTxs)
	return collectedTxs, err
}
