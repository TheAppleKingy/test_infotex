package database

import (
	"log/slog"
	"os"
	"test_infotex/ent"

	_ "github.com/lib/pq"
)

func NewClient() *ent.Client {
	database_url := os.Getenv("DATABASE_URL")
	if database_url == "" {
		slog.Error("database url was not provide")
		os.Exit(1)
	}
	client, err := ent.Open("postgres", database_url)
	if err != nil {
		slog.Error("error database connection", "error", err)
		os.Exit(1)
	}
	return client
}
