package model

type TxVolume struct {
	Datetime string `json:"datetime"`
	Denome   string `json:"denom"`
	Value    string `json:"value"`
}
