package handler

import "fmt"

type CreateProdutoRequest struct {
	Nome       string  `json:"nome"`
	Valor      float64 `json:"valor"`
	Quantidade int64   `json:"quantidade"`
	Descricao  string  `json:"descricao"`
}
type UpdateProdutoRequest struct {
	Nome       string  `json:"nome"`
	Valor      float64 `json:"valor"`
	Quantidade int64   `json:"quantidade"`
	Descricao  string  `json:"descricao"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (r *CreateProdutoRequest) Validate() error {
	if r.Nome == "" && r.Descricao == "" && r.Quantidade <= 0 && r.Valor <= 0 {
		return fmt.Errorf("request body is empty")
	}
	if r.Nome == "" {
		return errParamIsRequired("Nome", "string")
	}
	if r.Descricao == "" {
		return errParamIsRequired("Descricao", "string")
	}
	if r.Quantidade <= 0 {
		return errParamIsRequired("Quantidade", "string")
	}
	if r.Valor <= 0 {
		return errParamIsRequired("Valor", "string")
	}

	return nil
}
func (r *UpdateProdutoRequest) Validate() error {
	if r.Nome == "" && r.Descricao == "" && r.Quantidade <= 0 && r.Valor <= 0 {
		return fmt.Errorf("a least one field must be provied")
	}

	return nil
}
