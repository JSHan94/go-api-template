package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	glib "github.com/initia-labs/initia-apis/lib"
)

func GetBlockLatestQuery(c *gin.Context) (string, error) {
	return `{
		"size" : 1,
		"sort": [{"block.header.height": {"order": "desc"}}]
	}`, nil
}

func GetBlockByHeightQuery(c *gin.Context) (string, error) {
	type GetBlockByHeightPathParameter struct {
		Height uint64 `form:"height" uri:"height" binding:"required,gte=1"`
	}

	uri := &GetBlockByHeightPathParameter{}
	if err := glib.ValidateRequestParameters(c, &uri, nil); err != nil {
		return "", err
	}

	return fmt.Sprintf(`{
		"query": {
			"match" : {
				"block.header.height" : %d
			}
		}
	}`, uri.Height), nil
}

func GetBlockByTimeQuery(c *gin.Context) (string, error) {
	type GetBlockByTimePathParameter struct {
		Time string `uri:"time" binding:"required"`
	}

	uri := &GetBlockByTimePathParameter{}

	if err := glib.ValidateRequestParameters(c, &uri, nil); err != nil {
		return "", err
	}

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
	}`, uri.Time), nil
}

func GetBlockByHashQuery(c *gin.Context) (string, error) {
	type GetBlockByHashPathParameter struct {
		Hash string `uri:"hash" binding:"required"`
	}
	uri := &GetBlockByHashPathParameter{}

	if err := glib.ValidateRequestParameters(c, &uri, nil); err != nil {
		return "", err
	}

	return fmt.Sprintf(`{
		"query": {
			"match": {"block_id.hash": "%s"}
		}
	}`, uri.Hash), nil
}

func GetBlockAvgTimeQuery(c *gin.Context) (string, error) {
	type GetBlockAvgTimePathParameter struct {
		Height uint64 `uri:"height" binding:"required,gte=2"`
	}

	uri := &GetBlockAvgTimePathParameter{}

	if err := glib.ValidateRequestParameters(c, &uri, nil); err != nil {
		return "", err
	}

	return fmt.Sprintf(`{
		"query": {
			"range" : {
				"block.header.height" : {
					"lte" : %d
				}
			}
		},
		"size" : 1000,
		"sort": [{"block.header.height": {"order": "asc"}}]
	}`, uri.Height), nil
}

func GetBlocksQuery(c *gin.Context) (string, error) {
	type GetBlocksPathParameter struct {
		From uint64 `uri:"from" binding:"required,gte=1"`
		To   uint64 `uri:"to" binding:"required,gte=1"`
	}
	type GetBlocksQueryParameter struct {
		Order string `form:"order"`
	}

	uri := &GetBlocksPathParameter{}
	query := &GetBlocksQueryParameter{}

	if err := glib.ValidateRequestParameters(c, &uri, &query); err != nil {
		return "", err
	}

	return fmt.Sprintf(`{
		"query": {
			"range" : {
				"block.header.height" : {
					"gte" : %d,
					"lte" : %d
				}
			}
		},
		"sort" : [
			{"block.header.height" : {"order" : "%s"}}
	   	]
	}`, uri.From, uri.To, c.DefaultQuery("order", "desc")), nil
}
