package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log/slog"
	"os"
	"test_infotex/ent"
	"test_infotex/internal/infra"
	"time"

	_ "github.com/lib/pq"
)

func generateAddress() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		slog.Error("error creating address for wallets")
		os.Exit(1)
	}
	return hex.EncodeToString(bytes)
}

// This code create db schemas, apply migrations and create 10 wallets with random addresses and 100 balance value
func main() {
	client := infra.NewClient()
	defer client.Close()
	startConnectingTime := time.Now()
	for {
		if time.Since(startConnectingTime) > 6*time.Second {
			slog.Error("migrations could not apply")
			os.Exit(1)
		}
		err := client.Schema.Create(context.Background())
		if err == nil {
			break
		}
		slog.Error("error apply migrations", "error", err, "status", "retry")
		time.Sleep(1 * time.Second)
	}
	_, err := client.Wallet.Query().Limit(1).Only(context.Background())
	if ent.IsNotFound(err) {
		walletBuilders := make([]*ent.WalletCreate, 0, 10)
		for i := 0; i < 10; i++ {
			walletBuilders = append(walletBuilders, client.Wallet.Create().
				SetAddress(generateAddress()).
				SetBalance(10000))
		}
		if _, err := client.Wallet.CreateBulk(walletBuilders...).Save(context.Background()); err != nil {
			slog.Error("cannot create wallets", "error", err)
			os.Exit(1)
		}
	}
	slog.Info("migrations applied successfully")
}
