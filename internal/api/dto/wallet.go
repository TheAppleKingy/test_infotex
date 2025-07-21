package dto

type WalletBalance struct {
	Address string  `json:"address"`
	Balance float64 `json:"balance"`
}
