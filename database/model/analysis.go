package model

type TxVolume struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Denom string `json:"denom"`
	Value string `json:"value"`
}

type TxCount struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value int    `json:"value"`
}
