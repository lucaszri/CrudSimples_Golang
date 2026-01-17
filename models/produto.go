package models

type Produto struct {
	ID    int `json:"id"`
	Nome  string `json:"nome"`
	Valor float64 `json:"valor"`
}