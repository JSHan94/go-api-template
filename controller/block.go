package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gdatabase "github.com/initia-labs/initia-apis/database"
	gmodel "github.com/initia-labs/initia-apis/database/model"
	ghandler "github.com/initia-labs/initia-apis/handler"
)

var _ = &gmodel.CollectedBlock{}

// @Summary Get block
// @Description Get block information for a given height
// @Accept  json
// @Produce  json
// @Param  height  query  int  true  "height"
// @Success 200 {object} gmodel.CollectedBlock 	"block"
// @Router /v1/block [get]
func GetBlock(c *gin.Context) {
	var params ghandler.GetBlockQueryParameter
	if err := c.ShouldBindQuery(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	block, err := ghandler.GetBlock(gdatabase.IDX_BLOCK_BASIC, params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, block)
}

// @Summary Get blocks
// @Description Get blocks information for a given range
// @Accept  json
// @Produce  json
// @Param  from  query  integer  true  "from height"
// @Param  to  query  integer  true  "to height"
// @Success 200 {object} gmodel.CollectedBlocks 	"blocks"
// @Router /v1/blocks [get]
func GetBlocks(c *gin.Context) {
	var params ghandler.GetBlocksQueryParameter
	if err := c.ShouldBindQuery(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	blocks, err := ghandler.GetBlocks(gdatabase.IDX_BLOCK_BASIC, params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, blocks)
}
