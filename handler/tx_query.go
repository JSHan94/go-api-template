package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type GetTxQueryParameter struct {
	Hash string `form:"hash" binding:"required"`
}

type GetTxsQueryParameter struct {
	Limit string `form:"limit"`
	Order string `form:"order"`
}

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

	if params.Limit == "" {
		params.Limit = "10"
	}

	// allow only asc order
	return fmt.Sprintf(`{
		"query": {
			"match_all" : {}
		},
		"sort": [{"sequence": {"order": "asc"}}],
		"search_after": [%s],
		"size" : %s
	}`, offset, params.Limit), nil
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
