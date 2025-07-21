package main

import (
	"test_infotex/internal/api"
	"test_infotex/internal/infra"

	_ "github.com/lib/pq"
)

func main() {
	dbClient := infra.NewClient()
	router := api.GetRouter(dbClient)
	router.Run(":8080")
}
