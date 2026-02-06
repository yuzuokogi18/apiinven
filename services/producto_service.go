package services

import (
	"apiven/models"
	"apiven/repository"
	"errors"
)

type ProductoService struct {
	repo *repository.ProductoRepository
}

func NewProductoService() *ProductoService {
	return &ProductoService{
		repo: repository.NewProductoRepository(),
	}
}

func (s *ProductoService) Create(p *models.Producto, userID int) error {
	if p.Nombre == "" || p.Cantidad <= 0 {
		return errors.New("datos inválidos")
	}
	p.UsuarioID = userID
	return s.repo.Create(p)
}

func (s *ProductoService) GetAll(userID int) ([]models.Producto, error) {
	return s.repo.GetAllByUser(userID)
}
func (s *ProductoService) GetByID(id, userID int) (*models.Producto, error) {
	return s.repo.GetByID(id, userID)
}

func (s *ProductoService) Update(id int, p *models.Producto, userID int) error {
	if p.Nombre == "" || p.Cantidad <= 0 {
		return errors.New("datos inválidos")
	}
	return s.repo.Update(id, p, userID)
}

func (s *ProductoService) Delete(id, userID int) error {
	return s.repo.Delete(id, userID)
}
