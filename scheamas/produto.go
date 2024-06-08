package scheamas

import (
	"time"

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
	Id         uint      `json:"id"`
	CreateAt   time.Time `json:"createAt"`
	Nome       string    `json:"nome"`
	Valor      float64   `json:"valor"`
	Quantidade int64     `json:"quantidade"`
	Descricao  string    `json:"descricao"`
}
