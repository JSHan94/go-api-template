package model

type TxVolume struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Denom string `json:"denom"`
	Value int    `json:"value"`
}

type TxCount struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Value int    `json:"value"`
}
