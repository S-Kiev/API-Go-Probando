package handler

import (
	//"net/http"

	"github.com/S-Kiev/API-Go-Probando/middleware"
	"github.com/labstack/echo"
)

// RoutePerson .

func RoutePersona(e *echo.Echo, storage Storage) {
	h := newPersona(storage)
	person := e.Group("v1/personas")
	person.Use(middleware.Authentication)

	person.POST("", h.create)
	person.PUT("/:id", h.update)
	person.GET("", h.getAll)
	person.DELETE("/:id", h.delete)
	person.GET("/:id", h.getByID)
}

func RouteLogin(e *echo.Echo, storage Storage) {
	h := newLogin(storage)

	e.POST("/v1/login", h.login)
}
