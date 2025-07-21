package handlers

import (
	"net/http"
	"test_infotex/ent"
	"test_infotex/internal/api/dto"
	"test_infotex/internal/infra/repository"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	WalletRepo *repository.WalletRepo
}

func (h *WalletHandler) GetBalance(ctx *gin.Context) {
	address := ctx.Param("address")
	if address == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect wallet address",
		})
	}
	wallet, err := h.WalletRepo.GetByAddress(ctx.Request.Context(), address)
	if err != nil {
		if ent.IsNotFound(err) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "wallet with this address does not exist",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}
	respData := dto.WalletBalance{
		Address: wallet.Address,
		Balance: float64(wallet.Balance) / 100.0,
	}
	ctx.JSON(http.StatusOK, respData)
}
