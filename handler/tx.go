package handler

import (
	"fmt"

	gdatabase "github.com/initia-labs/initia-apis/database"
	gmodel "github.com/initia-labs/initia-apis/database/model"
	"github.com/sirupsen/logrus"
)

func GetTxByHash(indexName string, params GetTxQueryParameter) {
	// client := gdatabase.GetClient()

	// query := fmt.Sprintf(`{
	// 	"query": {
	// 		"bool" : {
	// 			"must" : {"match" : {"block.header.height" : %s}}
	// 		}
	// 	},
	// 	"_source": ["block_id", "block", "reward"]
	// }`, params.Height)

	logrus.Info("GetTxByHash called")
}

func GetTxsByAccount(indexName string, params GetTxsQueryParameter) (*gmodel.CollectedTxs, error) {
	logrus.Info("GetTxsByAccount called")
	query := fmt.Sprintf(`{
		"query": {
			"match": { 
				"tx.sender": "%s",
				"tx.chainid: "%s"
			}
		  },
		  "size": %s,
		  "sort": [ {"block.id": {"order": "desc"}} ]
	}`, params.Account, params.ChainId, params.Limit)

	return GetTxs(indexName, query)
}

func GetTxsByHeight(indexName string, params GetTxsQueryParameter) (*gmodel.CollectedTxs, error) {
	logrus.Info("GetTxsByHeight called")
	query := fmt.Sprintf(`{
		"query": {
			"match": { 
				"tx.sender": "%s",
				"tx.chainid: "%s"
			}
		  },
		  "size": %s,
		  "sort": [ {"block.id": {"order": "desc"}} ]
	}`, params.Height, params.ChainId, params.Limit)

	return GetTxs(indexName, query)
}

func GetTxsByOffset(indexName string, params GetTxsQueryParameter) (*gmodel.CollectedTxs, error) {
	logrus.Info("GetTxsByOffset called")

	query := fmt.Sprintf(`{
		"query": {
			"match": { 
				"tx.chainid: "%s"
			},
			range : { "block.id": { "gte": %s } },
		  },
		  "size": %s,
		  "sort": [ {"block.id": {"order": "desc"}} ]
	}`, params.Offset, params.ChainId, params.Limit)

	return GetTxs(indexName, query)
}

func GetTxs(indexName, query string) (*gmodel.CollectedTxs, error) {
	client := gdatabase.GetClient()

	hits, err := client.Search(indexName, query)
	if err != nil {
		logrus.Error("Error getting %s response: %s", query, err)
		return nil, err
	}

	collectedTxs := &gmodel.CollectedTxs{}
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"]
		collectedTx := &gmodel.CollectedTx{}
		if err := Decode(source.(map[string]interface{}), &collectedTx); err != nil {
			logrus.Error("Error decoding block", err)
			return nil, err
		}
		collectedTxs.Txs = append(collectedTxs.Txs, collectedTx)
	}

	return collectedTxs, nil
}
