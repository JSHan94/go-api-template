package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	gdatabase "github.com/initia-labs/initia-apis/database"
	gmodel "github.com/initia-labs/initia-apis/database/model"
	glib "github.com/initia-labs/initia-apis/lib"
)

func GetBlock(indexName string, query string) (*gmodel.CollectedBlock, error) {
	client := gdatabase.GetClient()

	hits, err := client.Search(indexName, query)
	if err != nil {
		return nil, err
	}
	collectedBlock := &gmodel.CollectedBlock{}
	err = glib.Decode(hits[0], collectedBlock)

	return collectedBlock, err
}

func GetBlocks(indexName string, query string) (*gmodel.CollectedBlocks, error) {
	client := gdatabase.GetClient()

	hits, err := client.Search(indexName, query)
	if err != nil {
		return nil, err
	}

	collectedBlocks := &gmodel.CollectedBlocks{}
	for _, hit := range hits {
		collectedBlock := &gmodel.CollectedBlock{}
		err = glib.Decode(hit, collectedBlock)
		if err != nil {
			return nil, err
		}
		collectedBlocks.Blocks = append(collectedBlocks.Blocks, collectedBlock)
	}

	return collectedBlocks, err
}

func GetBlockLatest(c *gin.Context, indexName string) (*gmodel.CollectedBlock, error) {
	query, err := GetBlockLatestQuery(c)
	if err != nil {
		return nil, err
	}

	return GetBlock(indexName, query)
}

func GetBlockByHeight(c *gin.Context, indexName string) (*gmodel.CollectedBlock, error) {
	query, err := GetBlockByHeightQuery(c)
	if err != nil {
		return nil, err
	}

	return GetBlock(indexName, query)
}

func GetBlockByTime(c *gin.Context, indexName string) (*gmodel.CollectedBlock, error) {
	query, err := GetBlockByTimeQuery(c)
	if err != nil {
		return nil, err
	}

	return GetBlock(indexName, query)
}

func GetBlockByHash(c *gin.Context, indexName string) (*gmodel.CollectedBlock, error) {
	query, err := GetBlockByHashQuery(c)
	if err != nil {
		return nil, err
	}

	return GetBlock(indexName, query)
}

func GetBlocksFromTo(c *gin.Context, indexName string) (*gmodel.CollectedBlocks, error) {
	query, err := GetBlocksQuery(c)
	if err != nil {
		return nil, err
	}
	collectedBlocks, err := GetBlocks(indexName, query)
	if err != nil {
		return nil, err
	}
	collectedBlocks.From = c.Param("from")
	collectedBlocks.To = c.Param("to")

	return collectedBlocks, nil
}

func GetBlockAvgTime(c *gin.Context, indexName string) (*gmodel.BlockAvgTime, error) {
	query, err := GetBlockAvgTimeQuery(c)
	if err != nil {
		return nil, err
	}
	collectedBlocks, err := GetBlocks(indexName, query)
	if err != nil {
		return nil, err
	}

	blockNum := len(collectedBlocks.Blocks) - 1
	if blockNum == 0 {
		return nil, errors.New("not enough blocks to calculate average block time")
	}

	startBlock := collectedBlocks.Blocks[0]
	endBlock := collectedBlocks.Blocks[blockNum]
	interval := endBlock.Block.Header.Time.Sub(startBlock.Block.Header.Time)
	avgBlockTime := float32(interval.Seconds()) / float32(blockNum)

	return &gmodel.BlockAvgTime{
		AvgTime: avgBlockTime,
	}, nil
}
