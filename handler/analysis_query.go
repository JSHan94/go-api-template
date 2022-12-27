package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type GetTxVolumeQueryParameter struct {
	From string `form:"from" binding:"required,gte=1"`
	To   string `form:"to" binding:"required,gte=1"`
	View string `form:"view"`
}

type GetTxCountQueryParameter struct {
	From string `form:"from" binding:"required,gte=1"`
	To   string `form:"to" binding:"required,gte=1"`
}

func GetTxVolumeQuery(c *gin.Context) (string, error) {
	params := &GetTxVolumeQueryParameter{}

	if err := c.ShouldBindQuery(&params); err != nil {
		return "", err
	}

	// TODO: define tx volume
	if params.View == "periodic" {
		return "", nil
	} else {
		return "", nil
	}
}

func GetTxCountQuery(c *gin.Context) (string, error) {
	params := &GetTxCountQueryParameter{}

	if err := c.ShouldBindQuery(&params); err != nil {
		return "", err
	}

	return fmt.Sprintf(`{
		"query": {
			"range": {
				"height": {
					"gte": %s,
					"lte": %s
				}
			}
		}
	}`, params.From, params.To), nil
}
