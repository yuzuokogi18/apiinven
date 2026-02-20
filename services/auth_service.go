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

/////////////////////////////
// REGISTER
/////////////////////////////

func (s *AuthService) Register(req *models.RegisterRequest) (*models.Usuario, string, error) {

	// Validación básica
	if req.Nombre == "" || req.Email == "" || req.Password == "" {
		return nil, "", errors.New("todos los campos son obligatorios")
	}

	// Verificar si ya existe el usuario
	existingUser, err := s.repo.GetByEmail(req.Email)
	if err == nil && existingUser != nil {
		return nil, "", errors.New("el usuario ya existe")
	}

	// Encriptar contraseña
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", errors.New("error al encriptar contraseña")
	}

	user := models.Usuario{
		Nombre:   req.Nombre,
		Email:    req.Email,
		Password: string(hash),
	}

	// Guardar usuario (esto debe asignar el ID)
	err = s.repo.Create(&user)
	if err != nil {
		return nil, "", err
	}

	// Generar token con ID real
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, "", errors.New("error al generar token")
	}

	// Ocultar password antes de devolver
	user.Password = ""

	return &user, token, nil
}

/////////////////////////////
// LOGIN
/////////////////////////////

func (s *AuthService) Login(req *models.LoginRequest) (*models.Usuario, string, error) {

	if req.Email == "" || req.Password == "" {
		return nil, "", errors.New("email y password son obligatorios")
	}

	user, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return nil, "", errors.New("usuario no encontrado")
	}

	// Verificar contraseña
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, "", errors.New("credenciales incorrectas")
	}

	// Generar token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, "", errors.New("error al generar token")
	}

	user.Password = ""

	return user, token, nil
}
