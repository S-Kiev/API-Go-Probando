package handler

import (
	"net/http"
	//"github.com/S-Kiev/API-Go-Probando/modelo"
)

// RoutePerson .
func RoutePersona(mux *http.ServeMux, storage Storage) {
	h := newPersona(storage)

	mux.HandleFunc("/v1/personas/create", h.create)
	mux.HandleFunc("/v1/personas/update", h.update)
	mux.HandleFunc("/v1/personas/get-all", h.getAll)
}
