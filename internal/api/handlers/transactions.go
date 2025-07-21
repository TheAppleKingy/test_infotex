package handlers

import (
	"net/http"
	"strconv"
	"test_infotex/internal/api/dto"
	"test_infotex/internal/application/service"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	TransactionService *service.TransactionService
}

func (th *TransactionHandler) Send(ctx *gin.Context) {
	var data dto.MakeTransactionDTO
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := th.TransactionService.TransferMoney(ctx.Request.Context(), data.From, data.To, data.Amount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (th *TransactionHandler) GetLast(ctx *gin.Context) {
	n := ctx.Query("count")
	var count int
	var err error
	if n != "" {
		count, err = strconv.Atoi(n)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "incorrect count param",
			})
			return
		}
	}
	transactionsLast, err := th.TransactionService.GetNLast(ctx.Request.Context(), count)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	respData := []dto.TransactionInfoDTO{}
	for _, transaction := range transactionsLast {
		respData = append(respData, dto.TransactionInfoDTO{
			From:   transaction.FromWallet,
			To:     transaction.ToWallet,
			Amount: transaction.Amount,
			At:     transaction.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, respData)
}
