package lib

import (
	"github.com/gin-gonic/gin"
)

func ValidateRequestParameters(c *gin.Context, uri interface{}, query interface{}) error {
	if uri != nil {
		if err := c.ShouldBindUri(uri); err != nil {
			return err
		}
	}
	if query != nil {
		if err := c.ShouldBindQuery(query); err != nil {
			return err
		}
	}
	return nil
}
