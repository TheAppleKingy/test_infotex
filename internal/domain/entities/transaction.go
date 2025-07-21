package entities

import "time"

type Transaction struct {
	ID         int
	Amount     float64
	FromWallet string
	ToWallet   string
	CreatedAt  time.Time
}
