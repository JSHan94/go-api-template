package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AbortWithStatusJSON(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": err.Error(),
	})
}
