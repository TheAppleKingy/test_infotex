package service

import (
	"context"
	"errors"
	"math"
	"test_infotex/internal/domain/entities"
	"test_infotex/internal/infra/repository"
)

type TransactionService struct {
	WalletRepo      *repository.WalletRepo
	TransactionRepo *repository.TransactionRepo
}

func (ts *TransactionService) convert(amount float64) (int, error) {
	converted := amount * 100
	if math.Round(converted) != converted {
		return 0, errors.New("monetary value can obly have 2 or 1 or 0 decimal places")
	}
	return int(converted), nil
}

func (ts *TransactionService) TransferMoney(ctx context.Context, fromAddr string, toAddr string, amount float64) error {
	if fromAddr == toAddr {
		return errors.New("cannot send money youself")
	}
	convertedAmount, err := ts.convert(amount)
	if err != nil {
		return err
	}
	fromWallet, err := ts.WalletRepo.GetByAddress(ctx, fromAddr)
	if err != nil {
		return err
	}
	toWallet, err := ts.WalletRepo.GetByAddress(ctx, toAddr)
	if err != nil {
		return err
	}
	if fromWallet.Balance < convertedAmount {
		return errors.New("not enough money have from wallet")
	}
	if err = ts.WalletRepo.UpdateBalance(ctx, fromAddr, fromWallet.Balance-convertedAmount); err != nil {
		return err
	}
	if err = ts.WalletRepo.UpdateBalance(ctx, toAddr, toWallet.Balance+convertedAmount); err != nil {
		return err
	}
	if err := ts.TransactionRepo.MakeTransaction(ctx, fromWallet.ID, toWallet.ID, convertedAmount); err != nil {
		return err
	}
	return nil
}

func (ts *TransactionService) GetNLast(ctx context.Context, n int) ([]entities.Transaction, error) {
	transactions, err := ts.TransactionRepo.GetNLast(ctx, n)
	if err != nil {
		return []entities.Transaction{}, err
	}
	return transactions, nil
}
