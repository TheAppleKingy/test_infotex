package repository

import (
	"test_infotex/ent"
	"test_infotex/ent/transaction"
	"test_infotex/internal/domain/entities"
	"time"

	"golang.org/x/net/context"
)

type TransactionRepo struct {
	client *ent.Client
}

func GetTransactionRepo(client *ent.Client) *TransactionRepo {
	return &TransactionRepo{
		client: client,
	}
}

func (tRepo *TransactionRepo) GetByTimeRange(ctx context.Context, from_time time.Time, to_time time.Time) ([]entities.Transaction, error) {
	transactions, err := tRepo.client.Transaction.Query().Where(transaction.CreatedAtLTE(to_time), transaction.CreatedAtGTE(from_time)).WithFromWallet().WithToWallet().All(ctx)
	if err != nil {
		return []entities.Transaction{}, err
	}
	transactionsEntities := []entities.Transaction{}
	for _, t := range transactions {
		transactionsEntities = append(transactionsEntities, entities.Transaction{
			ID:          uint(t.ID),
			Amount:      t.Amount,
			From_wallet: uint(t.Edges.FromWallet.ID),
			To_wallet:   uint(t.Edges.ToWallet.ID),
		})
	}
	return transactionsEntities, nil
}
