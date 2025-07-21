package service

import (
	"test_infotex/internal/infra/repository"
)

func NewTransactionService(walletRepo *repository.WalletRepo, tRepo *repository.TransactionRepo) *TransactionService {
	return &TransactionService{
		WalletRepo:      walletRepo,
		TransactionRepo: tRepo,
	}
}
