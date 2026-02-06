package models

import "time"

type Usuario struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type Categoria struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	UsuarioID int    `json:"usuario_id"`
}

type Producto struct {
	ID               int       `json:"id"`
	Nombre           string    `json:"nombre"`
	Cantidad         int       `json:"cantidad"`
	FechaVencimiento time.Time `json:"fecha_vencimiento"`
	CategoriaID      int       `json:"categoria_id"`
	UsuarioID        int       `json:"usuario_id"`
}

/* DTOs */
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
