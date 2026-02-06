package services

import (
	"apiven/models"
	"apiven/repository"
	"apiven/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.UsuarioRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		repo: repository.NewUsuarioRepository(),
	}
}

func (s *AuthService) Register(req *models.RegisterRequest) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	user := models.Usuario{
		Nombre:   req.Nombre,
		Email:    req.Email,
		Password: string(hash),
	}

	return s.repo.Create(&user)
}

func (s *AuthService) Login(req *models.LoginRequest) (string, error) {
	user, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return "", errors.New("usuario no encontrado")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("credenciales incorrectas")
	}

	return utils.GenerateToken(user.ID)
}
