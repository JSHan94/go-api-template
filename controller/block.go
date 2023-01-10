package controller

import (
	"github.com/gin-gonic/gin"
	gconfig "github.com/initia-labs/initia-apis/config"
	gmodel "github.com/initia-labs/initia-apis/database/model"
	ghandler "github.com/initia-labs/initia-apis/handler"
	glib "github.com/initia-labs/initia-apis/lib"
)

var _ = &gmodel.CollectedBlock{}

// @Tags block
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

// @Tags block
// @Summary Get block by height
// @Description Get block information for a given height
// @Accept  json
// @Produce  json
// @Param  height  path  int  true  "block height" minimum(1)
// @Success 200 {object} gmodel.CollectedBlock 	"block"
// @Router /v1/block/height/{height} [get]
func GetBlockByHeight(c *gin.Context) {
	block, err := ghandler.GetBlockByHeight(c, gconfig.IDX_BLOCK_BASIC)
	glib.Render(c, block, err)
}

// @Tags block
// @Summary Get block by hash
// @Description Get block information for a given hash
// @Accept  json
// @Produce  json
// @Param  hash  path  string  true  "block hash"
// @Success 200 {object} gmodel.CollectedBlock 	"block"
// @Router /v1/block/hash/{hash} [get]
func GetBlockByHash(c *gin.Context) {
	block, err := ghandler.GetBlockByHash(c, gconfig.IDX_BLOCK_BASIC)
	glib.Render(c, block, err)
}

// @Tags block
// @Summary Get block by time
// @Description Get a latest block information for a given time (yyyy-mm-ddTHH:MM:SS)
// @Accept  json
// @Produce  json
// @Param  time  path  string  true  "time" example(2022-11-08T19:47:10)
// @Success 200 {object} gmodel.CollectedBlock 	"block"
// @Router /v1/block/time/{time} [get]
func GetBlockByTime(c *gin.Context) {
	block, err := ghandler.GetBlockByTime(c, gconfig.IDX_BLOCK_BASIC)
	glib.Render(c, block, err)
}

// @Tags block
// @Summary Get average blocktime
// @Description Get average time of the previous 1000 blocks. If you give block height as 4500, it will give average time of blocks 3500 to 4500.
// @Accept  json
// @Produce  json
// @Param  height  path  integer  true  "block height" minimum(2)
// @Success 200 {object} gmodel.BlockAvgTime "average blocktime"
// @Router /v1/block/avgtime/{height} [get]
func GetBlockAvgTime(c *gin.Context) {
	blockAvgTime, err := ghandler.GetBlockAvgTime(c, gconfig.IDX_BLOCK_BASIC)
	glib.Render(c, blockAvgTime, err)
}

// @Tags block
// @Summary Get blocks
// @Description Get blocks information for a given range
// @Accept  json
// @Produce  json
// @Param  from  path  integer  true  "start height"
// @Param  to  path  integer  true  "end height"
// @Param  limit  query  int  false  "a limited number of records" default(10)
// @param  order query string  false "desc or asc" default(desc)
// @Success 200 {object} gmodel.CollectedBlocks 	"blocks"
// @Router /v1/blocks/{from}/{to} [get]
func GetBlocksFromTo(c *gin.Context) {
	blocks, err := ghandler.GetBlocksFromTo(c, gconfig.IDX_BLOCK_BASIC)
	glib.Render(c, blocks, err)
}
