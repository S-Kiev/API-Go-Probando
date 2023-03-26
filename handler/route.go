package handler

import (
	"net/http"

	"github.com/S-Kiev/API-Go-Probando/middleware"
)

// RoutePerson .
func RoutePersona(mux *http.ServeMux, storage Storage) {
	h := newPersona(storage)

	mux.HandleFunc("/v1/personas/create", middleware.Log(middleware.Authentication(h.create)))
	mux.HandleFunc("/v1/personas/update", h.update)
	mux.HandleFunc("/v1/personas/get-all", middleware.Log(h.getAll))
	mux.HandleFunc("/v1/personas/delete", middleware.Log(h.delete))
	mux.HandleFunc("/v1/persons/get-by-id", middleware.Log(h.getByID))
}
