package handler

import "github.com/S-Kiev/API-Go-Probando/modelo"

// Storage .
type Storage interface {
	Create(person *modelo.Persona) error
	Update(ID int, person *modelo.Persona) error
	Delete(ID int) error
	GetByID(ID int) (modelo.Persona, error)
	GetAll() (modelo.Personas, error)
}
