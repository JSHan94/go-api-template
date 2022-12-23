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
// @Success 200 {object} gmodel.CollectedBlock 	"block"
// @Router /v1/block/height/:height [get]
func GetBlockByHeight(c *gin.Context) {
	block, err := ghandler.GetBlockByHeight(c, gconfig.IDX_BLOCK_BASIC)
	if err != nil {
		AbortWithStatusJSON(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, block)
}

// @Summary Get average blocktime
// @Description Get average time of the previous 1000 blocks
// @Accept  json
// @Produce  json
// @Param  height  query  int  true  "current block height (should be greater than 2)"
// @Success 200 {object} gmodel.BlockAvgTime 	"average blocktime"
// @Router /v1/block/avgtime [get]
func GetBlockAvgTime(c *gin.Context) {
	blockAvgTime, err := ghandler.GetBlockAvgTime(c, gconfig.IDX_BLOCK_BASIC)
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
// @param  order query string  false "desc or asc (default: desc)"
// @Success 200 {object} gmodel.CollectedBlocks 	"blocks"
// @Router /v1/blocks [get]
func GetBlocks(c *gin.Context) {

	blocks, err := ghandler.GetBlocks(c, gconfig.IDX_BLOCK_BASIC)
	if err != nil {
		AbortWithStatusJSON(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, blocks)
}
