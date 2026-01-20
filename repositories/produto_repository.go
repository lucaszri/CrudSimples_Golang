package repositories

import (
	"database/sql"
	"obra-crud/models"
)

type ProdutoRepository struct {
	DB *sql.DB
}

func NewProdutoRepository(db *sql.DB) *ProdutoRepository {
	return &ProdutoRepository{DB: db}
}

func (r* ProdutoRepository) Create(produto models.Produto) (int64, error) {
	query := "INSERT INTO produtos (nome, valor) VALUES (?, ?)"
	result, err := r.DB.Exec(query, produto.Nome, produto.Valor)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r* ProdutoRepository) GetAll() ([]models.Produto, error) {
	query := "SELECT id, nome, valor FROM produtos"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produtos []models.Produto
	for rows.Next() {
		var produto models.Produto
		err := rows.Scan(&produto.ID, &produto.Nome, &produto.Valor)
		if err != nil {
			return nil, err
		}
		produtos = append(produtos, produto)
	}
	return produtos, nil
}

func (r* ProdutoRepository) GetByID(id int) (*models.Produto, error) {
	query := "SELECT id, nome, valor FROM produtos WHERE id = ?"
	row := r.DB.QueryRow(query, id)
	var produto models.Produto
	err := row.Scan(&produto.ID, &produto.Nome, &produto.Valor)
	if err != nil {
		return nil, err
	}
	return &produto, nil
}

func (r* ProdutoRepository) Update(produto models.Produto) error {
	query := "UPDATE produtos SET nome = ?, valor = ? WHERE id = ?"
	_, err := r.DB.Exec(query, produto.Nome, produto.Valor, produto.ID)
	return err
}

func (r* ProdutoRepository) Delete(id int) error {
	query := "DELETE FROM produtos WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}