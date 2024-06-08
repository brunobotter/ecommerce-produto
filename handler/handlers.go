package handler

import (
	"net/http"

	"github.com/brunobotter/ecommerce-produto/service"
	"github.com/brunobotter/ecommerce-produto/vo"
	"github.com/gin-gonic/gin"
)

func CreateProdutoHandler(ctx *gin.Context) {
	var request vo.CreateProdutoRequest
	if err := ctx.Bind(&request); err != nil {
		vo.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	produto, err := service.CreateProduto(request)
	if err != nil {
		vo.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	vo.SendSuccess(ctx, "Produto criado com sucesso", produto)
}

func DeleteProdutoHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		vo.SendError(ctx, http.StatusBadRequest, "ID é obrigatório")
		return
	}

	err := service.DeleteProduto(id)
	if err != nil {
		vo.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	vo.SendSuccess(ctx, "Produto deletado com sucesso", nil)
}

func ShowProdutoHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		vo.SendError(ctx, http.StatusBadRequest, "ID é obrigatório")
		return
	}

	produto, err := service.ShowProduto(id)
	if err != nil {
		vo.SendError(ctx, http.StatusNotFound, "Produto não encontrado")
		return
	}

	vo.SendSuccess(ctx, "Produto encontrado", produto)
}

func ListProdutoHandler(ctx *gin.Context) {
	produtos, err := service.ListProdutos()
	if err != nil {
		vo.SendError(ctx, http.StatusInternalServerError, "Erro ao listar produtos")
		return
	}

	vo.SendSuccess(ctx, "Produtos listados com sucesso", produtos)
}

func UpdateProdutoHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var request vo.UpdateProdutoRequest
	if err := ctx.BindJSON(&request); err != nil {
		vo.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	produto, err := service.UpdateProduto(id, request)
	if err != nil {
		vo.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	vo.SendSuccess(ctx, "Produto atualizado com sucesso", produto)
}

func VendaProdutoHandler(ctx *gin.Context) {

	var request vo.VendaProdutoRequest
	if err := ctx.BindJSON(&request); err != nil {
		vo.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	produto, err := service.VendaProduto(request)
	if err != nil {
		vo.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	vo.SendSuccess(ctx, "Produto vendido com sucesso", produto)
}
