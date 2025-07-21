package repository

import (
	"context"
	"test_infotex/ent"
	"test_infotex/ent/wallet"
	"test_infotex/internal/domain/entities"
)

type WalletRepo struct {
	Client *ent.Client
}

func (walletRepo *WalletRepo) GetByAddress(ctx context.Context, addr string) (entities.Wallet, error) {
	wallet, err := walletRepo.Client.Wallet.Query().Where(wallet.AddressEQ(addr)).Only(ctx)
	if err != nil {
		return entities.Wallet{}, err
	}
	return entities.Wallet{
		ID:      wallet.ID,
		Address: wallet.Address,
		Balance: wallet.Balance,
	}, nil
}

func (walletRepo *WalletRepo) UpdateBalance(ctx context.Context, forAddr string, to int) error {
	return walletRepo.Client.Wallet.Update().Where(wallet.AddressEQ(forAddr)).SetBalance(to).Exec(ctx)
}
