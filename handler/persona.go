package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/S-Kiev/API-Go-Probando/modelo"
)

type persona struct {
	storage Storage
}

func newPersona(storage Storage) persona {
	return persona{storage}
}

func (p *persona) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "Metodo no permitido"}`))
		return
	}

	data := modelo.Persona{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "La persona ..."}`))
		return
	}

	err = p.storage.Create(&data)
	if err != nil {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type": "error", "message": "Hubo un error al crear la persona"}`))
		return
	}

	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message_type": "error", "message": "Persona creada correctamente"}`))
	/*
		if r.Method != http.MethodPost {
			response := newResponse(Error, "Método no permitido", nil)
			responseJSON(w, http.StatusBadRequest, response)
			return
		}

		data := modelo.Persona{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
			responseJSON(w, http.StatusInternalServerError, response)
			return
		}

		err = p.storage.Create(&data)
		if err != nil {
			response := newResponse(Error, "Hubo un problema al crear la persona", nil)
			responseJSON(w, http.StatusInternalServerError, response)
			return
		}

		response := newResponse(Message, "Persona creada correctamente", nil)
		responseJSON(w, http.StatusCreated, response)
	*/
}

func (p *persona) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "Metodo no permitido"}`))
		return
	}

	Id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "El ID debe ser un numero entero positivo"}`))
		return
	}

	data := modelo.Persona{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "La persona no existe"}`))
		return
	}

	err = p.storage.Update(Id, &data)

	if err != nil {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type": "error", "message": "Hubo un error al actualizar el registro"}`))
		return
	}

	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message_type": "error", "message": "Persona actualizada correctamente"}`))

	/*
		if r.Method != http.MethodPut {
			response := newResponse(Error, "Método no permitido", nil)
			responseJSON(w, http.StatusBadRequest, response)
			return
		}

		ID, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			response := newResponse(Error, "El id debe ser un número entero positivo", nil)
			responseJSON(w, http.StatusBadRequest, response)
			return
		}

		data := modela.Persona{}
		err = json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
			responseJSON(w, http.StatusBadRequest, response)
			return
		}

		err = p.storage.Update(ID, &data)
		if err != nil {
			response := newResponse(Error, "Hubo un problema al obtener todas las personas", nil)
			responseJSON(w, http.StatusInternalServerError, response)
			return
		}

		response := newResponse(Message, "Persona actualizada correctamente", nil)
		responseJSON(w, http.StatusOK, response)
	*/
}

func (p *persona) delete(w http.ResponseWriter, r *http.Request) {
	/*
		if r.Method != http.MethodDelete {
			response := newResponse(Error, "Método no permitido", nil)
			responseJSON(w, http.StatusBadRequest, response)
			return
		}

		ID, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			response := newResponse(Error, "El id debe ser un número entero positivo", nil)
			responseJSON(w, http.StatusBadRequest, response)
			return
		}

		err = p.storage.Delete(ID)
		if errors.Is(err, modelo.ErrIDPersonaNoExiste) {
			response := newResponse(Error, "El ID de la persona no existe", nil)
			responseJSON(w, http.StatusBadRequest, response)
			return
		}
		if err != nil {
			response := newResponse(Error, "Ocurrió un error al elminar el registro", nil)
			responseJSON(w, http.StatusInternalServerError, response)
			return
		}

		response := newResponse(Message, "Ok", nil)
		responseJSON(w, http.StatusOK, response)
	*/
}

func (p *persona) getByID(w http.ResponseWriter, r *http.Request) {
	/*
		if r.Method != http.MethodGet {
			response := newResponse(Error, "Método no permitido", nil)
			responseJSON(w, http.StatusBadRequest, response)
			return
		}

		ID, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			response := newResponse(Error, "El id debe ser un número entero positivo", nil)
			responseJSON(w, http.StatusBadRequest, response)
			return
		}

		data, err := p.storage.GetByID(ID)
		if errors.Is(err, modelo.ErrIDPersonaNoExiste) {
			response := newResponse(Error, "El ID de la persona no existe", nil)
			responseJSON(w, http.StatusBadRequest, response)
			return
		}
		if err != nil {
			response := newResponse(Error, "Ocurrió un error al elminar el registro", nil)
			responseJSON(w, http.StatusInternalServerError, response)
			return
		}

		response := newResponse(Message, "Ok", data)
		responseJSON(w, http.StatusOK, response)
	*/
}

func (p *persona) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "Metodo no permitido"}`))
		return
	}

	resp, err := p.storage.GetAll()
	if err != nil {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type": "error", "message": "Hubo un error al recuperar todas las personas"}`))
		return
	}

	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type": "error", "message": "Hubo un error al convertir el slice en JSON"}`))
		return
	}

	/*
		if r.Method != http.MethodGet {
			response := newResponse(Error, "Método no permitido", nil)
			responseJSON(w, http.StatusBadRequest, response)
			return
		}

		data, err := p.storage.GetAll()


		if err != nil {
			response := newResponse(Error, "Hubo un problema al obtener todas las personas", nil)
			responseJSON(w, http.StatusInternalServerError, response)
			return
		}

		response := newResponse(Message, "Ok", data)
		responseJSON(w, http.StatusOK, response)
	*/
}
