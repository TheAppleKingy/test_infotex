package repository

import (
	"test_infotex/ent"
	"test_infotex/internal/domain/entities"

	"golang.org/x/net/context"
)

type TransactionRepo struct {
	Client *ent.Client
}

func (tr *TransactionRepo) GetNLast(ctx context.Context, n int) ([]entities.Transaction, error) {
	transactions, err := tr.Client.Transaction.Query().Order(ent.Desc("created_at")).Limit(n).WithFromWallet().WithToWallet().All(ctx)
	if err != nil {
		return []entities.Transaction{}, err
	}
	transactionsEntities := []entities.Transaction{}
	for _, t := range transactions {
		transactionsEntities = append(transactionsEntities, entities.Transaction{
			ID:         t.ID,
			Amount:     float64(t.Amount) / 100,
			FromWallet: t.Edges.FromWallet.Address,
			ToWallet:   t.Edges.ToWallet.Address,
			CreatedAt:  t.CreatedAt,
		})
	}
	return transactionsEntities, nil
}

func (tr *TransactionRepo) MakeTransaction(ctx context.Context, fromWallet int, toWallet int, amount int) error {
	if _, err := tr.Client.Transaction.Create().SetFromWalletID(fromWallet).SetToWalletID(toWallet).SetAmount(amount).Save(ctx); err != nil {
		return err
	}
	return nil
}
