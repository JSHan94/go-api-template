package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type GetBlocksQueryParameter struct {
	From  string `form:"from" binding:"required,gte=1"`
	To    string `form:"to" binding:"required,gte=1"`
	Order string `form:"order"`
}

type GetBlockAvgTimeQueryParameter struct {
	Height string `form:"height" binding:"gte=2"`
}

// Transaction params
type GetTxQueryParameter struct {
	Hash string `form:"hash" binding:"required"`
}

type GetTxsQueryParameter struct {
	Limit string `form:"limit"`
	Order string `form:"order"`
}

// block.go queries
func GetBlockByHeightQuery(c *gin.Context) (string, error) {
	height := c.Param("height")

	return fmt.Sprintf(`{
		"query": {
			"bool" : {
				"must" : {"match" : {"block.header.height" : %s}}
			}
		},
		"_source": ["block_id", "block"]
	}`, height), nil
}

func GetBlockAvgTimeQuery(c *gin.Context) (string, error) {
	var params GetBlockAvgTimeQueryParameter
	if err := c.ShouldBindQuery(&params); err != nil {
		return "", err
	}

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
	}`, params.Height), nil
}

func GetBlocksQuery(c *gin.Context) (string, error) {
	params := &GetBlocksQueryParameter{}

	if err := c.ShouldBindQuery(&params); err != nil {
		return "", err
	}
	if params.Order == "" {
		params.Order = "desc"
	}

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
			{"block.header.height" : {"order" : "%s"}}
	   	],	
		"_source": ["block_id", "block"]
	}`, params.From, params.To, params.Order), nil
}

// tx.go queries
func GetTxByHashQuery(c *gin.Context) (string, error) {
	hash := c.Param("hash")
	return fmt.Sprintf(`{
		"query": {
			"match" : {
				"txhash": "%s"
			}
		}
	}`, hash), nil
}

func GetTxsByOffsetQuery(c *gin.Context) (string, error) {
	offset := c.Param("offset")
	params := &GetTxsQueryParameter{}
	if err := c.ShouldBindQuery(&params); err != nil {
		return "", err
	}

	if params.Order == "" {
		params.Order = "desc"
	}

	if params.Limit == "" {
		params.Limit = "10"
	}

	return fmt.Sprintf(`{
		"query": {
			"match_all" : {}
		},
		"sort": [{"sequence": {"order": "%s"}}],
		"search_after": [%s],
		"size" : %s
	}`, params.Order, offset, params.Limit), nil
}

func GetTxsByAccountQuery(c *gin.Context) (string, error) {
	account := c.Param("account")
	params := &GetTxsQueryParameter{}
	if err := c.ShouldBindQuery(&params); err != nil {
		return "", err
	}

	if params.Order == "" {
		params.Order = "desc"
	}

	if params.Limit == "" {
		params.Limit = "10"
	}

	return fmt.Sprintf(`{
		"query": {
			"multi_match" : {
				"query" : "%s",
				"fields" : ["*"]
			}
		},
		"sort": [{"sequence": {"order": "%s"}}],
		"size" : %s
	}`, account, params.Order, params.Limit), nil
}

func GetTxsByHeightQuery(c *gin.Context) (string, error) {
	height := c.Param("height")
	params := &GetTxsQueryParameter{}
	if err := c.ShouldBindQuery(&params); err != nil {
		return "", err
	}

	if params.Order == "" {
		params.Order = "desc"
	}

	if params.Limit == "" {
		params.Limit = "10"
	}

	return fmt.Sprintf(`{
		"query": {
			"match" : {
				"height": "%s"
			}
		},
		"sort": [{"sequence": {"order": "%s"}}],
		"size": %s
	}`, height, params.Order, params.Limit), nil
}
