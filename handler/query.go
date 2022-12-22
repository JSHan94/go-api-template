package handler

import (
	"fmt"
)

// Block params
type GetBlockQueryParameter struct {
	Height string `form:"height" binding:"required,gte=1"`
}

type GetBlocksQueryParameter struct {
	From string `form:"from" binding:"required,gte=1"`
	To   string `form:"to" binding:"required,gte=1"`
}

type GetBlockAvgTimeQueryParameter struct {
	Height string `form:"height,gte=1"`
}

// Transaction params
type GetTxQueryParameter struct {
	Hash string `form:"hash" binding:"required"`
}

type GetTxsQueryParameter struct {
	Account string `form:"account"`
	Height  string `form:"height"`
	ChainId string `form:"chainid"`
	Offset  string `form:"offset"`
	Limit   string `form:"limit"`
}

// block.go queries
func GetBlockQuery(params GetBlockQueryParameter) string {
	return fmt.Sprintf(`{
		"query": {
			"bool" : {
				"must" : {"match" : {"block.header.height" : %s}}
			}
		},
		"_source": ["block_id", "block"]
	}`, params.Height)
}

func GetBlockAvgTimeQuery(params GetBlockAvgTimeQueryParameter) string {
	return fmt.Sprintf(`{
		"query": {
			"range" : {
				"block.header.height" : {
					"lte" : %s
				}
			}
		},
		"size" : 1000,
		"sort": [{"block.header.height": {"order": "asc"}}],
		"_source": ["block_id", "block"]
	}`, params.Height)
}

func GetBlocksQuery(params GetBlocksQueryParameter) string {
	return fmt.Sprintf(`{
		"query": {
			"range" : {
				"block.header.height" : {
					"gte" : %s,
					"lte" : %s
				}
			}
		},
		"sort" : [
			{"block.header.height" : {"order" : "desc"}}
	   	],	
		"_source": ["block_id", "block"]
	}`, params.From, params.To)
}

// tx.go queries
func GetTxQuery(params GetTxQueryParameter) string {
	return fmt.Sprintf(`{
		"query": {
			"match" : {
				"txhash": "%s"
			}
		}
	})`, params.Hash)
}

func GetTxsByOffsetQuery(params GetTxsQueryParameter) string {
	return fmt.Sprintf(`{
		"query": {
			"match_all" : {}
		},
		"sort": [{"height": {"order": "desc"}} , {"order" : {"order": "desc"}}],
		"search_after": [%s],
		"size" : %s
	})`, params.Offset, params.Limit)
}

func GetTxsByAccountQuery(params GetTxsQueryParameter) string {
	return fmt.Sprintf(`{
		"query": {
		}
	})`)
}

func GetTxsByHeightQuery(params GetTxsQueryParameter) string {
	return fmt.Sprintf(`{
		"query": {
			"range" : {
				"height" : {
					"lte" : %s
				}
			}
		},
		"size": %s,
		"sort": [{"height": {"order": "desc"}} , {"order" : {"order": "desc"}}]
	})`, params.Height, params.Limit)
}
