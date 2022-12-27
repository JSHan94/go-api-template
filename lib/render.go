package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, res interface{}, err error) {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}
