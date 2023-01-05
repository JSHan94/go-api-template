package model

type UnconfirmedTxsResult struct {
	JSONRPC string              `json:"jsonrpc"`
	ID      int                 `json:"id"`
	Result  MempoolTransactions `json:"result"`
}

type MempoolTransactions struct {
	ChainID    string   `json:"chain_id"`
	NumTxs     string   `json:"n_txs"`
	Total      string   `json:"total"`
	TotalBytes string   `json:"total_bytes"`
	Txs        []string `json:"txs"`
}

type GasPrice struct {
	Denom string `json:"denom"`
	Value string `json:"value"`
}
