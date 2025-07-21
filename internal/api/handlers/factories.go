package handlers

import (
	"test_infotex/ent"
	"test_infotex/internal/application/service"
	"test_infotex/internal/infra/repository"
)

func NewWalletHandler(client *ent.Client) *WalletHandler {
	walletRepo := repository.NewWalletRepo(client)
	return &WalletHandler{
		WalletRepo: walletRepo,
	}
}

func NewTransactionHandler(client *ent.Client) *TransactionHandler {
	walletRepo := repository.NewWalletRepo(client)
	tRepo := repository.NewTransactionRepo(client)
	tService := service.NewTransactionService(walletRepo, tRepo)
	return &TransactionHandler{
		TransactionService: tService,
	}
}
