package controller

import (
	"github.com/gin-gonic/gin"
	gconfig "github.com/initia-labs/initia-apis/config"
	gmodel "github.com/initia-labs/initia-apis/database/model"
	ghandler "github.com/initia-labs/initia-apis/handler"
	glib "github.com/initia-labs/initia-apis/lib"
)

var _ = &gmodel.CollectedBlock{}

// @Summary Get latest block
// @Description Get latest block
// @Accept  json
// @Produce  json
// @Success 200 {object} gmodel.CollectedBlock 	"block"
// @Router /v1/block/latest [get]
func GetBlockLatest(c *gin.Context) {
	block, err := ghandler.GetBlockLatest(c, gconfig.IDX_BLOCK_BASIC)
	glib.Render(c, block, err)
}

// @Summary Get block by time
// @Description Get block information for a given height
// @Accept  json
// @Produce  json
// @Success 200 {object} gmodel.CollectedBlock 	"block"
// @Router /v1/block/height/:height [get]
func GetBlockByHeight(c *gin.Context) {
	block, err := ghandler.GetBlockByHeight(c, gconfig.IDX_BLOCK_BASIC)
	glib.Render(c, block, err)
}

// @Summary Get block by hash
// @Description Get block information for a given hash
// @Accept  json
// @Produce  json
// @Success 200 {object} gmodel.CollectedBlock 	"block"
// @Router /v1/block/hash/:hash [get]
func GetBlockByHash(c *gin.Context) {
	block, err := ghandler.GetBlockByHash(c, gconfig.IDX_BLOCK_BASIC)
	glib.Render(c, block, err)
}

// @Summary Get block by time
// @Description Get a latest block information for a given time (e.g., 2022-11-08T02:47:10)
// @Accept  json
// @Produce  json
// @Success 200 {object} gmodel.CollectedBlock 	"block"
// @Router /v1/block/time/:time [get]
func GetBlockByTime(c *gin.Context) {
	block, err := ghandler.GetBlockByTime(c, gconfig.IDX_BLOCK_BASIC)
	glib.Render(c, block, err)
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
	glib.Render(c, blockAvgTime, err)
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
func GetBlocksFromTo(c *gin.Context) {
	blocks, err := ghandler.GetBlocksFromTo(c, gconfig.IDX_BLOCK_BASIC)
	glib.Render(c, blocks, err)
}
