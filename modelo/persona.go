package modelo

// Comunidad es la estructura de una comunidad
type Comunidad struct {
	// Nombre de una comunidad.
	Nombre string `json:"nombre"`
}

// Comunidades slice de comunidades
type Comunidades []Comunidad

// Persona estructura de una persona
type Persona struct {
	// Nombre de la persona Ej: Ezequiel
	Nombre string `json:"nombre"`
	// Edad de la persona Ej: 40
	Edad uint8 `json:"edad"`
	// Comunidades comunidades a las que pertenece una persona
	Comunidades Comunidades `json:"comunidades"`
}

// Personas slice de personas
type Personas []Persona
