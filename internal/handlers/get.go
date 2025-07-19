package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAny(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
