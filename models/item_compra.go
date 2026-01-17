package models

import "time"
	
type Item_compra struct {
	ID         int       `json:"id"`
	ProdutoID  int       `json:"produto_id"`
	Quantidade int       `json:"quantidade"`
	Data       time.Time `json:"data"`
	Produto    Produto   `json:"produto,omitempty"`
}