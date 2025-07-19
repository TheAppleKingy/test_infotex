package repository

import (
	"context"
	"test_infotex/ent"
	"test_infotex/ent/wallet"
	"test_infotex/internal/domain/entities"
)

type WalletRepo struct {
	client *ent.Client
}

func GetWalletRepo(client *ent.Client) *WalletRepo {
	return &WalletRepo{client: client}
}

func (walletRepo *WalletRepo) GetByAddress(ctx context.Context, addr string) (entities.Wallet, error) {
	wallet, err := walletRepo.client.Wallet.Query().Where(wallet.AddressEQ(addr)).Only(ctx)
	if err != nil {
		return entities.Wallet{}, err
	}
	return entities.Wallet{
		ID:      uint(wallet.ID),
		Address: wallet.Address,
		Balance: wallet.Balance,
	}, nil
}

func (walletRepo *WalletRepo) UpdateBalance(ctx context.Context, forAddr string, to uint) error {
	return walletRepo.client.Wallet.Update().Where(wallet.AddressEQ(forAddr)).SetBalance(to).Exec(ctx)
}
