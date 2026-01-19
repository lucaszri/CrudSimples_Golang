package app

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"github.com/joho/godotenv"
	"obra-crud/database"
	"obra-crud/handlers"
	"obra-crud/repositories"
)

type App struct {
	DB *database.Database
	Port string
}

func NewApp() (*App, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: arquivo .env não encontrado")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := database.NewDatabase(dsn)
	if err != nil {
		return nil, err
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &App{DB: db, Port: port}, nil
}

func (app *App) Run() error {
	produtoRepo := repositories.NewProdutoRepository(app.DB.GetDB())
	produtoHandler := handlers.NewProdutoHandler(produtoRepo)

	http.Handle("/produtos", produtoHandler)
	http.Handle("/produtos/", produtoHandler)

	log.Printf("Servidor rodando na porta %s", app.Port)

	return http.ListenAndServe(":" + app.Port, nil)
}

func (app *App) Close() error {
	if app.DB != nil {
		return app.DB.Close()
	}
	return nil
}