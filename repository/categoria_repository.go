package repository

import (
	"apiven/config"
	"apiven/models"
)

type CategoriaRepository struct{}

func NewCategoriaRepository() *CategoriaRepository {
	return &CategoriaRepository{}
}

func (r *CategoriaRepository) Create(c *models.Categoria) error {
	_, err := config.DB.Exec(
		"INSERT INTO categorias (nombre, usuario_id) VALUES (?, ?)",
		c.Nombre,
		c.UsuarioID,
	)
	return err
}
