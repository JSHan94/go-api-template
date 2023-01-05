package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	glib "github.com/initia-labs/initia-apis/lib"
)

func GetTxVolumeQuery(c *gin.Context) (string, error) {
	type GetTxVolumePathParameter struct {
		Start string `uri:"start" binding:"required"`
		End   string `uri:"end" binding:"required"`
	}
	uri := &GetTxVolumePathParameter{}

	if err := glib.ValidateRequestParameters(c, &uri, nil); err != nil {
		return "", err
	}

	return fmt.Sprintf(`{
		"query": {
			"range": {
				"timestamp": {
					"gte": "%s",
					"lt": "%s"
				}
			}
		}
	}`, uri.Start, uri.End), nil
}

func GetTxCountQuery(c *gin.Context) (string, error) {
	type GetTxCountPathParameter struct {
		Start string `uri:"start" binding:"required"`
		End   string `uri:"end" binding:"required"`
	}
	uri := &GetTxCountPathParameter{}
	if err := glib.ValidateRequestParameters(c, &uri, nil); err != nil {
		return "", err
	}

	return fmt.Sprintf(`{
		"query": {
			"range": {
				"timestamp": {
					"gt": "%s",
					"lt": "%s"
				}
			}
		}
	}`, uri.Start, uri.End), nil
}
