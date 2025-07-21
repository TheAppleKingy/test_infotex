package dto

import "time"

type MakeTransactionDTO struct {
	From   string  `json:"from" binding:"required"`
	To     string  `json:"to" binding:"required"`
	Amount float64 `json:"amount" binding:"required,gte=0"`
}

type TransactionInfoDTO struct {
	From   string    `json:"from"`
	To     string    `json:"to"`
	Amount float64   `json:"amount"`
	At     time.Time `json:"at"`
}
