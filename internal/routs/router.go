package routs

import (
	"test_infotex/internal/handlers"
	"test_infotex/internal/handlers/middleware"

	"github.com/gin-gonic/gin"
)

// GetRouter returns a fully-fledged *gin.Engine that only needs to be started.
func GetRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", handlers.GetAny)
	router.POST("/", handlers.GetAny)
	router.Use(middleware.MethodNotAllowedMiddleware(router.Routes()))
	return router
}
