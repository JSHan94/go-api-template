package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gconfig "github.com/initia-labs/initia-apis/config"
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
		AbortWithStatusJSON(c, err)
		return
	}
	block, err := ghandler.GetBlock(gconfig.IDX_BLOCK_BASIC, params)
	if err != nil {
		AbortWithStatusJSON(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, block)
}

// @Summary Get average blocktime
// @Description Get average time of the last 1000 blocks
// @Accept  json
// @Produce  json
// @Param  height  query  int  false  "request height"
// @Success 200 {object} gmodel.BlockAvgTime 	"average blocktime"
// @Router /v1/block/avgtime [get]
func GetBlockAvgTime(c *gin.Context) {
	var params ghandler.GetBlockAvgTimeQueryParameter
	if err := c.ShouldBindQuery(&params); err != nil {
		AbortWithStatusJSON(c, err)
		return
	}

	blockAvgTime, err := ghandler.GetBlockAvgTime(gconfig.IDX_BLOCK_BASIC, params)
	if err != nil {
		AbortWithStatusJSON(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, blockAvgTime)
}

// @Summary Get blocks
// @Description Get blocks information for a given range
// @Accept  json
// @Produce  json
// @Param  from  query  integer  true  "start height"
// @Param  to  query  integer  true  "end height"
// @Success 200 {object} gmodel.CollectedBlocks 	"blocks"
// @Router /v1/blocks [get]
func GetBlocks(c *gin.Context) {
	var params ghandler.GetBlocksQueryParameter
	if err := c.ShouldBindQuery(&params); err != nil {
		AbortWithStatusJSON(c, err)
		return
	}

	blocks, err := ghandler.GetBlocks(gconfig.IDX_BLOCK_BASIC, params)
	if err != nil {
		AbortWithStatusJSON(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, blocks)
}
