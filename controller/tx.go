package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gdatabase "github.com/initia-labs/initia-apis/database"
	gmodel "github.com/initia-labs/initia-apis/database/model"
	ghandler "github.com/initia-labs/initia-apis/handler"
	tm "github.com/tendermint/tendermint/types"
)

// type check gmodel.Transactions
var _ = &gmodel.CollectedTx{}
var _ = &gmodel.CollectedTxs{}
var _ = &tm.BlockID{}
var _ = &tm.Block{}

// @Summary Get txs
// @Description Get transactions matching with given params
// @Accept  json
// @Produce  json
// @Param  account	query	string	false	"Sender account address"
// @Param  block  query  int  false  "Block height"
// @Param  chainid  query  string  false  "Chain ID"
// @Param  limit  query  int  false  "Items per page (default: 10)"
// @Param  offset  query  int  false  "Offset"
// @Success 200 {object} gmodel.CollectedTxs "List of transactions"
// @Router /v1/txs [get]
func GetTxs(c *gin.Context) {
	var params ghandler.GetTxsQueryParameter
	if err := c.ShouldBindQuery(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if params == (ghandler.GetTxsQueryParameter{}) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "need to provide at least one query parameter (account or height or offset)",
		})
		return
	}

	txIndexName := gdatabase.IDX_TX_BASIC
	// check query.account is empty
	if params.Account != "" {
		ghandler.GetTxsByAccount(txIndexName, params)
	} else if params.Height != "" {
		ghandler.GetTxsByHeight(txIndexName, params)
	} else {
		ghandler.GetTxsByOffset(txIndexName, params)
	}

}

// @Summary Get gas price
// @Description Get minimum gas price
// @Accept  json
// @Produce  json
// @Success 200 {obejct} gmodel.GasPrice "uinit gas price"
// @Router /v1/txs/gasprice [get]
func GetGasPrice(c *gin.Context) {
	panic("not implemented yet")
}

// @Summary Get mempool transaction
// @Description Get mempool transaction with txhash
// @Accept  json
// @Produce  json
// @Param  txhash  query  string  true  "Transaction hash"
// @Success 200 {obejct} gmodel.MempoolTransaction "Mempool transaction"
// @Router /v1/mempool/tx [get]
func GetMempoolTx(c *gin.Context) {
	panic("not implemented yet")
}

// @Summary Get mempool transactions
// @Description Get mempool transactions
// @Accept  json
// @Produce  json
// @Param  account  query  string  false  "Account address"
// @Success 200 {obejct} gmodel.MempoolTransactions "Mempool transactions"
// @Router /v1/mempool/txs [get]
func GetMempoolTxs(c *gin.Context) {
	panic("not implemented yet")
}
