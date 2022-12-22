package handler

import (
	"errors"

	gdatabase "github.com/initia-labs/initia-apis/database"
	gmodel "github.com/initia-labs/initia-apis/database/model"
)

func GetBlock(indexName string, params GetBlockQueryParameter) (*gmodel.CollectedBlock, error) {
	client := gdatabase.GetClient()
	query := GetBlockQuery(params)

	hits, err := client.Search(indexName, query)
	if err != nil {
		return nil, err
	}

	collectedBlock := &gmodel.CollectedBlock{}
	err = DecodeOne(hits, collectedBlock)

	return collectedBlock, err
}

func GetBlockAvgTime(indexName string, params GetBlockAvgTimeQueryParameter) (*gmodel.BlockAvgTime, error) {
	client := gdatabase.GetClient()
	query := GetBlockAvgTimeQuery(params)

	hits, err := client.Search(indexName, query)
	if err != nil {
		return nil, err
	}

	blockNum := len(hits) - 1
	if blockNum == 0 {
		return nil, errors.New("not enough blocks to calculate average block time")
	}

	var collectedBlocks []gmodel.CollectedBlocks
	if err := DecodeMany(hits, &collectedBlocks); err != nil {
		return nil, err
	}

	startBlock := collectedBlocks[0]
	endBlock := collectedBlocks[blockNum]
	interval := endBlock.CollectedBlock.Block.Header.Time.Sub(startBlock.CollectedBlock.Block.Header.Time)
	avgBlockTime := float32(interval.Seconds()) / float32(blockNum)

	return &gmodel.BlockAvgTime{
		AvgTime: avgBlockTime,
	}, nil
}

func GetBlocks(indexName string, params GetBlocksQueryParameter) ([]gmodel.CollectedBlocks, error) {
	client := gdatabase.GetClient()
	query := GetBlocksQuery(params)

	hits, err := client.Search(indexName, query)
	if err != nil {
		return nil, err
	}

	var collectedBlocks []gmodel.CollectedBlocks
	err = DecodeMany(hits, &collectedBlocks)
	return collectedBlocks, err
}
