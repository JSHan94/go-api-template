package controller

import (
	"github.com/gin-gonic/gin"
	gconfig "github.com/initia-labs/initia-apis/config"
	gmodel "github.com/initia-labs/initia-apis/database/model"
	ghandler "github.com/initia-labs/initia-apis/handler"
	glib "github.com/initia-labs/initia-apis/lib"
	tm "github.com/tendermint/tendermint/types"
)

// type check gmodel.Transactions
var _ = &gmodel.CollectedTx{}
var _ = &gmodel.CollectedTxs{}
var _ = &tm.BlockID{}
var _ = &tm.Block{}

// @Summary Get a tx
// @Description Get a transaction matching with given hash
// @Accept  json
// @Produce  json
// @Param  hash	path  string	true	"transaction hash"
// @Success 200 {object} gmodel.CollectedTx "a transaction"
// @Router /v1/tx/{hash} [get]
func GetTxByHash(c *gin.Context) {
	tx, err := ghandler.GetTxByHash(c, gconfig.IDX_TX_BASIC)
	glib.Render(c, tx, err)
}

// @Summary Get txs
// @Description Get transactions related with given account address
// @Accept  json
// @Produce  json
// @Param  account  path  string true  "account address"
// @Param  limit  query  int  false  "a limited number of records" default(10000)
// @Param  order  query  string  false  "desc or asc" default(desc)
// @Success 200 {object} gmodel.CollectedTxs "List of transactions"
// @Router /v1/txs/account/{account} [get]
func GetTxsByAccount(c *gin.Context) {
	txs, err := ghandler.GetTxsByAccount(c, gconfig.IDX_TX_BASIC)
	glib.Render(c, txs, err)
}

// @Summary Get txs
// @Description Get transactions from given offset. If offset is 0, it returns all transactions.
// @Accept  json
// @Produce  json
// @Param  offset  path  int true  "offset to skip items"
// @Param  limit  query  int  false  "a limited number of records" default(10000)
// @Success 200 {object} gmodel.CollectedTxs "List of transactions"
// @Router /v1/txs/offset/{offset} [get]
func GetTxsByOffset(c *gin.Context) {
	txs, err := ghandler.GetTxsByOffset(c, gconfig.IDX_TX_BASIC)
	glib.Render(c, txs, err)
}

// @Summary Get txs
// @Description Get transactions from given block height
// @Accept  json
// @Produce  json
// @Param  height  path  int true  "block height" mininum(1)
// @Param  limit  query  int  false  "a limited number of records" default(10000)
// @Param  order  query  string  false  "desc or asc" default(desc)
// @Success 200 {object} gmodel.CollectedTxs "List of transactions"
// @Router /v1/txs/height/{height} [get]
func GetTxsByHeight(c *gin.Context) {
	txs, err := ghandler.GetTxsByHeight(c, gconfig.IDX_TX_BASIC)
	glib.Render(c, txs, err)
}

// @Summary Get gas price
// @Description Get minimum gas price
// @Accept  json
// @Produce  json
// @Success 200 {obejct} gmodel.GasPrice "uinit minimum gas price"
// @Router /v1/txs/gasprice [get]
func GetGasPrice(c *gin.Context) {
	gasPrice := gmodel.GasPrice{
		Denom: gconfig.NATIVE_TOKEN_SYMBOL,
		Value: gconfig.MIN_GAS_PRICE,
	}
	glib.Render(c, gasPrice, nil)
}

// // @Summary Get mempool transaction
// // @Description Get mempool transaction with txhash
// // @Accept  json
// // @Produce  json
// // @Param  txhash  query  string  true  "Transaction hash"
// // @Success 200 {obejct} gmodel.MempoolTransaction "Mempool transaction"
// // @Router /v1/tx/mempool [get]
// func GetMempoolTx(c *gin.Context) {
// 	panic("not implemented yet")
// }

// // @Summary Get mempool transactions
// // @Description Get mempool transactions
// // @Accept  json
// // @Produce  json
// // @Param  account  query  string  false  "Account address"
// // @Success 200 {obejct} gmodel.MempoolTransactions "Mempool transactions"
// // @Router /v1/txs/mempool [get]
// func GetMempoolTxs(c *gin.Context) {
// 	panic("not implemented yet")
// }
