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

func GetBlockLatestQuery(c *gin.Context) (string, error) {
	return `{
		"size" : 1,
		"sort": [{"block.header.height": {"order": "desc"}}]
	}`, nil
}

func GetBlockByHeightQuery(c *gin.Context) (string, error) {
	height := c.Param("height")

	return fmt.Sprintf(`{
		"query": {
			"match" : {
				"block.header.height" : %s
			}
		}
	}`, height), nil
}

func GetBlockByTimeQuery(c *gin.Context) (string, error) {
	height := c.Param("time") //"2022-11-07T16:39:17"

	return fmt.Sprintf(`{
		"query": {
			"range": {
			  "block.header.time": {
				"lt": "%s"
			  }
			}
		},
		"sort" : {"block.header.height" : "desc" },
		"size" : 1
	}`, height), nil
}

func GetBlockByHashQuery(c *gin.Context) (string, error) {
	hash := c.Param("hash")

	return fmt.Sprintf(`{
		"query": {
			"match": {"block_id.hash": "%s"}
		}
	}`, hash), nil
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
		"sort": [{"block.header.height": {"order": "asc"}}]
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
	   	]
	}`, params.From, params.To, params.Order), nil
}
