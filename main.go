package main

import (
	"fmt"
	"log"
	"net/http"
	"obra-crud/database"
	"obra-crud/handlers"
	"obra-crud/repositories"
)

func main() {
	database.Connect()
	defer database.Close()

	produtoRepo := repositories.NewProdutoRepository(database.DB)
	produtoHandler := handlers.NewProdutoHandler(produtoRepo)
	
	http.Handle("/produtos", produtoHandler)
	http.Handle("/produtos/", produtoHandler)

	fmt.Println("🚀 Servidor rodando na porta 8080")
	fmt.Println("📡 API disponível em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println("Hello, World!")
}