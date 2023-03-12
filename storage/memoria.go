package storage

//No tiene base de datos, funciona en memoria

import (
	"fmt"

	"github.com/S-Kiev/API-Go-Probando/modelo"
)

// Memoria .
type Memoria struct {
	ultimoID int
	Personas map[int]modelo.Persona
}

// NewMemoria .
func NewMemoria() Memoria {
	personas := make(map[int]modelo.Persona)

	return Memoria{
		ultimoID: 0,
		Personas: personas,
	}
}

// Crear .
func (m *Memoria) Create(persona *modelo.Persona) error {
	if persona == nil {
		return modelo.ErrPersonaNoPuedeSerNil
	}

	m.ultimoID++
	m.Personas[m.ultimoID] = *persona

	return nil
}

// Update actualiza una persona en el slice de memoria
func (m *Memoria) Update(ID int, persona *modelo.Persona) error {
	if persona == nil {
		return modelo.ErrPersonaNoPuedeSerNil
	}

	if _, ok := m.Personas[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, modelo.ErrIDPersonaNoExiste)
	}

	m.Personas[ID] = *persona

	return nil
}

// Delete borra de la memoria la persona
func (m *Memoria) Delete(ID int) error {
	if _, ok := m.Personas[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, modelo.ErrIDPersonaNoExiste)
	}

	delete(m.Personas, ID)

	return nil
}

// GetByID retorna una persona por el ID
func (m *Memoria) GetByID(ID int) (modelo.Persona, error) {
	persona, ok := m.Personas[ID]
	if !ok {
		return persona, fmt.Errorf("ID: %d: %w", ID, modelo.ErrIDPersonaNoExiste)
	}

	return persona, nil
}

// GetAll retorna todas las personas que están en la memoria
func (m *Memoria) GetAll() (modelo.Personas, error) {
	var result modelo.Personas
	for _, v := range m.Personas {
		result = append(result, v)
	}

	return result, nil
}
