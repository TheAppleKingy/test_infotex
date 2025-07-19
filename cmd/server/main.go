package main

import (
	"test_infotex/internal/routs"

	_ "github.com/lib/pq"
)

func main() {
	router := routs.GetRouter()
	router.Run(":8080")
}
