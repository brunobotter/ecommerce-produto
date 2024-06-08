package handler

import (
	"net/http"

	"github.com/brunobotter/ecommerce-produto/scheamas"
	"github.com/gin-gonic/gin"
)

func UpdateOpenningHandler(ctx *gin.Context) {
	request := UpdateProdutoRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Param("id")
	logger.Infof("id %s", id)
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}
	produto := scheamas.Produto{}
	if err := db.First(&produto, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "produto not found")
		return
	}
	if request.Nome != "" {
		produto.Nome = request.Nome
	}
	if request.Descricao != "" {
		produto.Descricao = request.Descricao
	}
	if request.Quantidade > 0 {
		produto.Quantidade = request.Quantidade
	}
	if request.Valor > 0 {
		produto.Valor = request.Valor
	}
	if err := db.Save(&produto).Error; err != nil {
		logger.Errorf("error updating produto %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	sendSuccess(ctx, "update-produto", produto)
}
