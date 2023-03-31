package main

import (
	"log"
	//"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/S-Kiev/API-Go-Probando/authorization"
	"github.com/S-Kiev/API-Go-Probando/handler"
	"github.com/S-Kiev/API-Go-Probando/storage"
)

func main() {

	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("no se pudo cargar los certificados: %v", err)
	}
	store := storage.NewMemoria()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	handler.RoutePersona(e, &store)
	handler.RouteLogin(e, &store)

	log.Println("Servidor iniciado en el puerto 8080")
	err = e.Start(":8080")
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}
