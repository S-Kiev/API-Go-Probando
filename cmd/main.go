package main

import (
	"log"
	"net/http"

	"github.com/S-Kiev/API-Go-Probando/handler"
	"github.com/S-Kiev/API-Go-Probando/storage"
)

func main() {
	store := storage.NewMemoria()
	mux := http.NewServeMux()

	handler.RoutePersona(mux, &store)

	log.Println("Servidor iniciado en el puerto 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}
