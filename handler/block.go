package handler

import (
	"fmt"

	gdatabase "github.com/initia-labs/initia-apis/database"
	gmodel "github.com/initia-labs/initia-apis/database/model"
	"github.com/sirupsen/logrus"
)

func GetBlock(indexName string, params GetBlockQueryParameter) (*gmodel.CollectedBlock, error) {
	client := gdatabase.GetClient()

	query := fmt.Sprintf(`{
		"query": {
			"bool" : {
				"must" : {"match" : {"block.header.height" : %s}}
			}
		},
		"_source": ["block_id", "block", "reward"]
	}`, params.Height)

	hits, err := client.Search(indexName, query)
	if err != nil {
		logrus.Error("Error getting %s response: %s", query, err)
		return nil, err
	}

	collectedBlock := &gmodel.CollectedBlock{}
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"]
		if err := Decode(source.(map[string]interface{}), &collectedBlock); err != nil {
			logrus.Error("Error decoding block", err)
			return nil, err
		}
	}

	return collectedBlock, nil
}

func GetBlocks(indexName string, params GetBlocksQueryParameter) (*gmodel.CollectedBlocks, error) {
	client := gdatabase.GetClient()

	// FIXME : https://guiyomi.tistory.com/18 (Field data type)
	query := fmt.Sprintf(`{
		"query": {
			"range" : {
				"block.header.height" : {
					"gte" : %s,
					"lte" : %s
				}
			}
		},
		"sort" : [
			{"block.header.heigh" : {"order" : "desc"}}
	   	],	
		"_source": ["block_id", "block", "reward"]
	}`, params.From, params.To)

	hits, err := client.Search(indexName, query)
	if err != nil {
		logrus.Error("Error getting %s response: %s", query, err)
		return nil, err
	}

	collectedBlocks := &gmodel.CollectedBlocks{}
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"]
		collectedBlock := &gmodel.CollectedBlock{}
		if err := Decode(source.(map[string]interface{}), &collectedBlock); err != nil {
			logrus.Error(err)
			return nil, err
		}
		collectedBlocks.Blocks = append(collectedBlocks.Blocks, collectedBlock)
	}

	return collectedBlocks, nil
}
