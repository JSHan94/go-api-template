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

// @Summary Get tx volume
// @Description Get transaction volume for a given range
// @Accept  json
// @Produce  json
// @Param  from	 query	string	true "from block height"
// @Param  to	 query	string	true "to block height"
// @Param  view	query	string	true	"Periodic or cumulative (default: cumulative)"
// @Success 200 {object} gmodel.TxVolume 	"tx volume"
// @Router /v1/analysis/tx-volume [get]
func GetTxVolume(c *gin.Context) {
	txVolume, err := ghandler.GetTxVolume(c, gconfig.IDX_EVENT_BASIC)
	glib.Render(c, txVolume, err)
}

// @Summary Get tx count
// @Description Get tx count for a given range
// @Accept  json
// @Produce  json
// @Param  from	 query	string	true "from block height"
// @Param  to	 query	string	true "to block height"
// @Success 200 {object} gmodel.TxCount "tx count"
// @Router /v1/analysis/tx-count [get]
func GetTxCount(c *gin.Context) {
	txCount, err := ghandler.GetTxCount(c, gconfig.IDX_TX_BASIC)
	glib.Render(c, txCount, err)
}
