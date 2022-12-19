package controller

import (
	"github.com/gin-gonic/gin"
	gmodel "github.com/initia-labs/initia-apis/database/model"
)

var _ = &gmodel.TxVolume{}

// @Summary Get tx volume
// @Description Get transaction volume
// @Accept  json
// @Produce  json
// @Param  denom query	string	true	"Denomination"
// @Param  view	query	string	true	"Choose periodic or cumulative (default: cumulative)"
// @Success 200 {object} gmodel.TxVolume 	"tx volume"
// @Router /v1/dashboard/tx/volume [get]
func GetTxVolume(c *gin.Context) {
	panic("not implemented yet")
}
