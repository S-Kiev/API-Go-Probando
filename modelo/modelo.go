package modelo

import "errors"

var (
	// ErrPersonaNoPuedeSerNil la persona no puede ser nula
	ErrPersonaNoPuedeSerNil = errors.New("la persona no puede ser nula")
	// ErrIDPersonaNoExiste la persona no existe
	ErrIDPersonaNoExiste = errors.New("la persona no existe")
)
