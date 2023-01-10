package controller

import (
	"github.com/gin-gonic/gin"
	gconfig "github.com/initia-labs/initia-apis/config"
	gmodel "github.com/initia-labs/initia-apis/database/model"
	ghandler "github.com/initia-labs/initia-apis/handler"
	glib "github.com/initia-labs/initia-apis/lib"
)

var _ = &gmodel.TxVolume{}
var _ = &gmodel.TxCount{}

// @Tags analysis
// @Summary Get tx volume
// @Description Get transaction volume for a given time range
// @Accept  json
// @Produce  json
// @Param  start  path	string	true "start time" example(2022-11-08T19:47:10)
// @Param  end	 path	string	true "end time" example(2022-11-18T20:47:10)
// @Success 200 {object} gmodel.TxVolume 	"tx volume"
// @Router /v1/analysis/tx-volume/{start}/{end} [get]
func GetTxVolume(c *gin.Context) {
	txVolume, err := ghandler.GetTxVolume(c, gconfig.IDX_EVENT_BASIC)
	glib.Render(c, txVolume, err)
}

// @Tags analysis
// @Summary Get tx count
// @Description Get tx count for a given time range
// @Accept  json
// @Produce  json
// @Param  start  path	 string	true "start time" example(2022-11-08T19:47:10)
// @Param  end	 path	string	true "end time" example(2022-11-18T20:47:10)
// @Success 200 {object} gmodel.TxCount "tx count"
// @Router /v1/analysis/tx-count/{start}/{end} [get]
func GetTxCount(c *gin.Context) {
	txCount, err := ghandler.GetTxCount(c, gconfig.IDX_TX_BASIC)
	glib.Render(c, txCount, err)
}
