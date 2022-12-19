package handler

type GetBlockQueryParameter struct {
	Height string `form:"height" binding:"required"`
}

type GetBlocksQueryParameter struct {
	From string `form:"from" binding:"required"`
	To   string `form:"to" binding:"required"`
}

type GetTxQueryParameter struct {
	Hash string `form:"hash" binding:"required"`
}

type GetTxsQueryParameter struct {
	Account string `form:"account"`
	Height  string `form:"height"`
	ChainId string `form:"chainid"`
	Offset  string `form:"offset"`
	Limit   string `form:"limit"`
}
