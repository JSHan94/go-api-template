package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	glib "github.com/initia-labs/initia-apis/lib"
)

type GetTxsQueryParameter struct {
	Limit string `form:"limit"`
	Order string `form:"order"`
}

func GetTxByHashQuery(c *gin.Context) (string, error) {
	type GetTxByHashPathParameter struct {
		Hash string `uri:"hash" binding:"required"`
	}

	uri := &GetTxByHashPathParameter{}

	if err := glib.ValidateRequestParameters(c, &uri, nil); err != nil {
		return "", err
	}

	return fmt.Sprintf(`{
		"query": {
			"match" : {
				"txhash": "%s"
			}
		}
	}`, uri.Hash), nil
}

func GetTxsByOffsetQuery(c *gin.Context) (string, error) {
	// to use gte=0, don't use required binding
	// https://github.com/gin-gonic/gin/issues/1246
	type GetTxsByOffsetPathParameter struct {
		Offset int64 `uri:"offset" binding:"gte=0"`
	}

	uri := &GetTxsByOffsetPathParameter{}
	query := &GetTxsQueryParameter{}

	if err := glib.ValidateRequestParameters(c, &uri, &query); err != nil {
		return "", err
	}

	// allow only asc order
	return fmt.Sprintf(`{
		"query": {
			"match_all" : {}
		},
		"sort": [{"sequence": {"order": "asc"}}],
		"search_after": [%d],
		"size" : %s
	}`, uri.Offset, c.DefaultQuery("limit", "10000")), nil
}

func GetTxsByAccountQuery(c *gin.Context) (string, error) {
	type GetTxsByAccountPathParameter struct {
		Account string `uri:"account" binding:"required"`
	}
	uri := &GetTxsByAccountPathParameter{}
	query := &GetTxsQueryParameter{}
	if err := glib.ValidateRequestParameters(c, &uri, &query); err != nil {
		return "", err
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
	}`, uri.Account, c.DefaultQuery("order", "desc"), c.DefaultQuery("limit", "10000")), nil
}

func GetTxsByHeightQuery(c *gin.Context) (string, error) {
	type GetTxsByHeightPathParameter struct {
		Height uint64 `uri:"height" binding:"required,gte=1"`
	}
	uri := &GetTxsByHeightPathParameter{}
	query := &GetTxsQueryParameter{}
	if err := glib.ValidateRequestParameters(c, &uri, &query); err != nil {
		return "", err
	}

	return fmt.Sprintf(`{
		"query": {
			"match" : {
				"height": %d
			}
		},
		"sort": [{"sequence": {"order": "%s"}}],
		"size": %s
	}`, uri.Height, c.DefaultQuery("order", "desc"), c.DefaultQuery("limit", "10000")), nil
}
