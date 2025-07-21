package api

import (
	"test_infotex/ent"
	"test_infotex/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

// GetRouter returns a fully-fledged *gin.Engine that only needs to be started.
func GetRouter(client *ent.Client) *gin.Engine {
	walletHandler := handlers.NewWalletHandler(client)
	transactionHandler := handlers.NewTransactionHandler(client)
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	router.GET("/api/wallet/:address/balance", walletHandler.GetBalance)
	router.GET("/api/transactions", transactionHandler.GetLast)
	router.POST("/api/send", transactionHandler.Send)
	return router
}
