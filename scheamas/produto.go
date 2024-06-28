package scheamas

import (
	"gorm.io/gorm"
)

type Produto struct {
	gorm.Model
	Nome       string
	Valor      float64
	Quantidade int64
	Descricao  string
}

type ProdutoResponse struct {
	Id         uint    `json:"id"`
	Nome       string  `json:"nome"`
	Valor      float64 `json:"valor"`
	Quantidade int64   `json:"quantidade"`
	Descricao  string  `json:"descricao"`
}

func ToprodutoResponse(produto Produto) ProdutoResponse {
	return ProdutoResponse{
		Id:         produto.ID,
		Nome:       produto.Nome,
		Valor:      produto.Valor,
		Quantidade: produto.Quantidade,
		Descricao:  produto.Descricao,
	}
}
