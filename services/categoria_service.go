package services

import (
	"apiven/models"
	"apiven/repository"
)

type CategoriaService struct {
	repo *repository.CategoriaRepository
}

func NewCategoriaService() *CategoriaService {
	return &CategoriaService{
		repo: repository.NewCategoriaRepository(),
	}
}

func (s *CategoriaService) Create(c *models.Categoria, userID int) error {
	c.UsuarioID = userID
	return s.repo.Create(c)
}
