package repository

import (
	"apiven/config"
	"apiven/models"
)

type UsuarioRepository struct{}

func NewUsuarioRepository() *UsuarioRepository {
	return &UsuarioRepository{}
}

func (r *UsuarioRepository) Create(u *models.Usuario) error {
	query := `
	INSERT INTO usuarios (nombre, email, password_hash)
	VALUES (?, ?, ?)
	`
	_, err := config.DB.Exec(query, u.Nombre, u.Email, u.Password)
	return err
}

func (r *UsuarioRepository) GetByEmail(email string) (*models.Usuario, error) {
	row := config.DB.QueryRow(
		"SELECT id, nombre, email, password_hash FROM usuarios WHERE email = ?",
		email,
	)

	var u models.Usuario
	err := row.Scan(&u.ID, &u.Nombre, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
