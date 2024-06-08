package service

import (
	"fmt"

	"github.com/brunobotter/ecommerce-produto/scheamas"
	"github.com/brunobotter/ecommerce-produto/vo"
	"gorm.io/gorm"
)

// db é a instância do banco de dados global
var db *gorm.DB

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

func ShowProduto(id string) (scheamas.Produto, error) {
	produto := scheamas.Produto{}
	if err := db.First(&produto, id).Error; err != nil {
		return produto, err
	}

	return produto, nil
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

func VendaProduto(venda vo.VendaProdutoRequest) (scheamas.Produto, error) {
	produto := scheamas.Produto{}
	if err := db.First(&produto, venda.Id).Error; err != nil {
		return produto, err
	}
	if produto.Quantidade >= venda.Quantidade {
		venda := produto.Quantidade - venda.Quantidade
		produto.Quantidade = venda
	} else {
		return produto, fmt.Errorf("erro ao deletar produto com ID %s", venda.Id)
	}
	if err := db.Save(&produto).Error; err != nil {
		return produto, err
	}
	return produto, nil
}
