package repositories

import (
	"database/sql"
	"obra-crud/models"
	"time"
)

type ItemCompraRepository struct {
	DB *sql.DB
}

func NewItemCompraRepository(db *sql.DB) *ItemCompraRepository {
	return &ItemCompraRepository{DB: db}
}

func (r *ItemCompraRepository) Create(itemCompra models.Item_compra) (int64, error) {
	query := "INSERT INTO itens_compra (produto_id, quantidade, data) VALUES (?, ?, ?)"
	result, err := r.DB.Exec(query, itemCompra.ProdutoID, itemCompra.Quantidade, time.Now())
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *ItemCompraRepository) GetAll() ([]models.Item_compra, error) {
	query := `SELECT ic.id, ic.produto_id, ic.quantidade, ic.data,
			  p.id, p.nome, p.valor
			  FROM itens_compra ic
			  JOIN produtos p ON ic.produto_id = p.id`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var itens []models.Item_compra
	for rows.Next() {
		var item models.Item_compra
		var produto models.Produto
		err := rows.Scan(&item.ID, &item.ProdutoID, &item.Quantidade, &item.Data,
			&produto.ID, &produto.Nome, &produto.Valor)
		if err != nil {
			return nil, err
		}
		item.Produto = produto
		itens = append(itens, item)
	}
	return itens, nil
}

func (r *ItemCompraRepository) Delete(id int) (*models.Item_compra, error) {
	var item models.Item_compra
	query := `SELECT ic.id, ic.produto_id, ic.quantidade, ic.data,
			  p.id, p.nome, p.valor
			  FROM itens_compra ic
			  JOIN produtos p ON ic.produto_id = p.id
			  WHERE ic.id = ?`
	err := r.DB.QueryRow(query, id).Scan(&item.ID, &item.ProdutoID, &item.Quantidade, &item.Data,
		&item.Produto.ID, &item.Produto.Nome, &item.Produto.Valor)
	if err != nil {
		return nil, err
	}
	query = "DELETE FROM itens_compra WHERE id = ?"
	_, err = r.DB.Exec(query, id)
	if err != nil {
		return nil, err
	}
	return &item, nil
}