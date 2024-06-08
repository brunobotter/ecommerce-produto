package handler

import (
	"net/http"

	"github.com/brunobotter/ecommerce-produto/scheamas"
	"github.com/gin-gonic/gin"
)

func ListOpenningHandler(ctx *gin.Context) {
	produto := []scheamas.Produto{}
	if err := db.Find(&produto).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing produtos!")
	}
	sendSuccess(ctx, "list-produtos", produto)
}
