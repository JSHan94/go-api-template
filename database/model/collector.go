package model

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tm "github.com/tendermint/tendermint/types"
)

// FIXME: use archiver indexer/collector/types.go
type BlockReward struct {
	Reward           string
	Commission       string
	RewardPerVal     string
	CommissionPerVal string
}

type CollectedBlock struct {
	BlockID *tm.BlockID `json:"block_id" swaggertype:"object"`
	Block   *tm.Block   `json:"block" swaggertype:"object"`
}

type CollectedBlocks struct {
	From   string            `json:"from"`
	To     string            `json:"to"`
	Blocks []*CollectedBlock `json:"blocks"`
}

type CollectedTxs struct {
	Limit  string         `json:"limit"`
	Offset string         `json:"offset"`
	Txs    []*CollectedTx `json:"txs"`
}

type CollectedTx struct {
	Tx         json.RawMessage `json:"tx" swaggertype:"string"`
	TxResponse json.RawMessage `json:"tx_response" swaggertype:"string"`
	Code       uint32          `json:"code"`
	Codespace  string          `json:"codespace"`
	GasUsed    int64           `json:"gas_used"`
	GasWanted  int64           `json:"gas_wanted"`
	Height     int64           `json:"height"`
	RawLog     string          `json:"raw_log"`
	Logs       json.RawMessage `json:"logs" swaggertype:"string"`
	TxHash     string          `json:"txhash"`
	Timestamp  time.Time       `json:"timestamp" swaggertype:"string"`
	Index      uint            `json:"index"`    // tx order in a block
	Sequence   uint64          `json:"sequence"` // sequence of whole txs, not sender's sequence
}

type CollectedEvents struct {
	Type       string               `json:"type"`
	Attributes []CollectedAttribute `json:"attributes"`
	Height     int64                `json:"height"`
	Timestamp  time.Time            `json:"timestamp"`
}

type CollectedAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Index bool   `json:"index"`
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

type CollectedAppendix struct {
	indexBytes      []byte             // internal use only
	StakingRatio    sdk.Dec            `json:"staking_ratio"`
	Issuances       sdk.Coins          `json:"issuances"`
	CommunityPool   sdk.DecCoins       `json:"community_pool"`
	BondedTokens    sdk.Coin           `json:"bonded_tokens"`
	NotBondedTokens sdk.Coin           `json:"not_bonded_tokens"`
	Reward          BlockReward        `json:"reward"`
	Network         []CollectedNetwork `json:"network"`
	Height          int64              `json:"height"`
	Timestamp       time.Time          `json:"timestamp"`
}
