package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)
type Database struct {
	DB *sql.DB
}

func NewDatabase(dsn string) (*Database, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	fmt.Println("Conectado ao MySQL com sucesso!")

	return &Database{DB: db}, nil
}

func (d *Database) Close() error {
	if d.DB != nil {
		return d.DB.Close()
	}
	return nil
}

func (d *Database) GetDB() *sql.DB {
	return d.DB
}
