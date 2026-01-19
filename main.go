package main

import (
	"log"
	"obra-crud/app"
)

func main() {
 application, err := app.NewApp()
 if err != nil {
  log.Fatalf("Erro ao iniciar a aplicação: %v", err)
 }
 defer application.Close()

 if err := application.Run(); err != nil {
  log.Fatalf("Erro ao rodar a aplicação: %v", err)
 }
}