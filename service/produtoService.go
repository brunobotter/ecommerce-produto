package service

import (
	"fmt"

	"github.com/brunobotter/ecommerce-produto/config"
	"github.com/brunobotter/ecommerce-produto/scheamas"
	"github.com/brunobotter/ecommerce-produto/vo"
	"gorm.io/gorm"
)

// db é a instância do banco de dados global
var db *gorm.DB

func InitializeService(database *gorm.DB) {
	db = database
}
func CreateProduto(request vo.CreateProdutoRequest) (scheamas.Produto, error) {
	produto := scheamas.Produto{
		Nome:       request.Nome,
		Valor:      request.Valor,
		Quantidade: request.Quantidade,
		Descricao:  request.Descricao,
	}

	if err := db.Create(&produto).Error; err != nil {
		return produto, err
	}

	return produto, nil
}

func DeleteProduto(id string) error {
	produto := scheamas.Produto{}
	if err := db.First(&produto, id).Error; err != nil {
		return fmt.Errorf("produto com ID %s não encontrado", id)
	}

	if err := db.Delete(&produto).Error; err != nil {
		return fmt.Errorf("erro ao deletar produto com ID %s", id)
	}

	return nil
}

func ShowProduto(id string) (scheamas.ProdutoResponse, error) {
	logger := config.GetLogger("ShowProduto")
	produto := scheamas.Produto{}
	if err := db.First(&produto, id).Error; err != nil {
		return scheamas.ProdutoResponse{}, err
	}
	responseProduto := scheamas.ToprodutoResponse(produto)
	logger.Debugf("produto: %v", responseProduto)
	return responseProduto, nil
}

func ListProdutos() ([]scheamas.Produto, error) {
	var produtos []scheamas.Produto
	if err := db.Find(&produtos).Error; err != nil {
		return nil, err
	}

	return produtos, nil
}

func UpdateProduto(id string, request vo.UpdateProdutoRequest) (scheamas.Produto, error) {
	produto := scheamas.Produto{}
	if err := db.First(&produto, id).Error; err != nil {
		return produto, err
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
		return produto, err
	}

	return produto, nil
}

func VendaProduto(venda vo.VendaProdutoRequest) (vo.VendaResponse, error) {
	logger := config.GetLogger("VendaProduto")
	produto := scheamas.Produto{}
	var vendas vo.VendaResponse
	if err := db.First(&produto, venda.Id).Error; err != nil {
		return vo.VendaResponse{}, err
	}
	if (produto.Quantidade >= venda.Quantidade) && venda.Quantidade > 0 {
		quantidade := produto.Quantidade - venda.Quantidade
		produto.Quantidade = quantidade
		total := produto.Valor * float64(venda.Quantidade)
		vendas = VendaResponse(produto, venda.Quantidade, total)
	} else {
		return vo.VendaResponse{}, fmt.Errorf("erro quantidade de itens insuficiente no estoque!")
	}
	if err := db.Save(&produto).Error; err != nil {
		return vo.VendaResponse{}, err
	}
	logger.Debugf("vendas %v", vendas)
	return vendas, nil
}

func VendaResponse(produto scheamas.Produto, quantidade int64, total float64) vo.VendaResponse {
	return vo.VendaResponse{
		Id:         produto.ID,
		Nome:       produto.Nome,
		Quantidade: quantidade,
		Valor:      total,
		Descricao:  produto.Descricao,
	}
}
