package main

import (
	"context"
	"log/slog"
	"os"
	"test_infotex/ent"

	_ "github.com/lib/pq"
)

func main() {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		slog.Error("database url was not provide")
		os.Exit(1)
	}
	client, err := ent.Open("postgres", databaseUrl)
	if err != nil {
		slog.Error("error database connection", "error", err)
		os.Exit(1)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		slog.Error("error apply migrations", "error", err)
		os.Exit(1)
	}
	slog.Info("migrations applied successfully")
}
