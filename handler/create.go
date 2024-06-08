package handler

import (
	"net/http"

	"github.com/brunobotter/ecommerce-produto/scheamas"
	"github.com/gin-gonic/gin"
)

func CreateOpenningHandler(ctx *gin.Context) {
	request := CreateProdutoRequest{}
	ctx.Bind(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	produto := scheamas.Produto{
		Nome:       request.Nome,
		Valor:      request.Valor,
		Quantidade: request.Quantidade,
		Descricao:  request.Descricao,
	}

	logger.Infof("request recieved: %v", request)
	if err := db.Create(&produto).Error; err != nil {
		logger.Errorf("Error creating produto: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "Create produto", produto)
}
