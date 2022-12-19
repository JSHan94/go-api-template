package model

import (
	"encoding/json"
	"time"

	abci "github.com/tendermint/tendermint/abci/types"
	tm "github.com/tendermint/tendermint/types"
)

type BlockReward struct {
	Reward           string
	Commission       string
	RewardPerVal     string
	CommissionPerVal string
}

type CollectedBlock struct {
	BlockID *tm.BlockID `json:"block_id" swaggertype:"object"`
	Block   *tm.Block   `json:"block" swaggertype:"object"`
	Reward  BlockReward `json:"reward"`
}

type CollectedBlocks struct {
	Blocks []*CollectedBlock `json:"blocks"`
}

type CollectedTxs struct {
	Txs []*CollectedTx `json:"txs"`
}

type CollectedTx struct {
	Tx         json.RawMessage `json:"tx"`
	TxResponse json.RawMessage `json:"tx_response"`
	Code       uint32          `json:"code"`
	Codespace  string          `json:"codespace"`
	GasUsed    int64           `json:"gas_used"`
	GasWanted  int64           `json:"gas_wanted"`
	Height     int64           `json:"height"`
	RawLog     string          `json:"raw_log"`
	Logs       json.RawMessage `json:"logs"`
	TxHash     string          `json:"txhash"`
	Timestamp  time.Time       `json:"timestamp"`
	Sender     string          `json:"sender"`
}

type CollectedEvents struct {
	BeginBlock *abci.ResponseBeginBlock  `json:"begin_block"`
	EndBlock   *abci.ResponseEndBlock    `json:"end_block"`
	DeliverTxs []*abci.ResponseDeliverTx `json:"deliver_txs"`
}

type CollectedNetwork struct {
	Denom     string `json:"denom"`
	Supply    string `json:"supply"`
	MarketCap string `json:"market_cap"`
	TxVolume  string `json:"tx_volume"`
}

type CollectedGeneral struct {
	StakingRatio    float64           `json:""`
	Issuances       map[string]string `json:""`
	CommunityPool   map[string]string `json:""`
	BondedTokens    string            `json:""`
	NotBondedTokens string            `json:""`
}

type TxRecord struct {
	Tx         json.RawMessage `json:"tx"`
	TxResponse json.RawMessage `json:"tx_response"`
}
