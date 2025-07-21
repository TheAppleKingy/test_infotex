package repository

import (
	"test_infotex/ent"
)

func NewWalletRepo(client *ent.Client) *WalletRepo {
	return &WalletRepo{
		Client: client,
	}
}

func NewTransactionRepo(client *ent.Client) *TransactionRepo {
	return &TransactionRepo{
		Client: client,
	}
}
