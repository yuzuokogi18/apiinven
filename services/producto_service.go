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

// 🔹 Crear producto
func (s *ProductoService) Create(p *models.Producto, userID int) error {
	if p == nil {
		return errors.New("producto vacío")
	}

	if p.Nombre == "" {
		return errors.New("el nombre es obligatorio")
	}

	if p.Cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor a 0")
	}

	// asignar usuario desde el token
	p.UsuarioID = userID

	return s.repo.Create(p)
}

// 🔹 Obtener todos los productos del usuario
func (s *ProductoService) GetAll(userID int) ([]models.Producto, error) {
	return s.repo.GetAllByUser(userID)
}

// 🔹 Obtener producto por ID
func (s *ProductoService) GetByID(id, userID int) (*models.Producto, error) {
	if id <= 0 {
		return nil, errors.New("id inválido")
	}
	return s.repo.GetByID(id, userID)
}

// 🔹 Actualizar producto
func (s *ProductoService) Update(id int, p *models.Producto, userID int) error {
	if id <= 0 {
		return errors.New("id inválido")
	}

	if p == nil {
		return errors.New("producto vacío")
	}

	if p.Nombre == "" {
		return errors.New("el nombre es obligatorio")
	}

	if p.Cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor a 0")
	}

	return s.repo.Update(id, p, userID)
}

// 🔹 Eliminar producto
func (s *ProductoService) Delete(id, userID int) error {
	if id <= 0 {
		return errors.New("id inválido")
	}
	return s.repo.Delete(id, userID)
}
