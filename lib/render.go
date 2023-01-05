package lib

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, res interface{}, err error) {
	if err != nil {
		switch {
		case errors.Is(err, ErrNotFound):
			c.IndentedJSON(http.StatusOK, res)
		default:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
	} else {
		c.IndentedJSON(http.StatusOK, res)
	}
}
