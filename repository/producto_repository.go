package repository

import (
	"apiven/config"
	"apiven/models"
)

type ProductoRepository struct{}

func NewProductoRepository() *ProductoRepository {
	return &ProductoRepository{}
}

func (r *ProductoRepository) Create(p *models.Producto) error {
	query := `
	INSERT INTO productos (nombre, cantidad, fecha_vencimiento, categoria_id, usuario_id)
	VALUES (?, ?, ?, ?, ?)
	`
	_, err := config.DB.Exec(
		query,
		p.Nombre,
		p.Cantidad,
		p.FechaVencimiento,
		p.CategoriaID,
		p.UsuarioID,
	)
	return err
}

func (r *ProductoRepository) GetAllByUser(userID int) ([]models.Producto, error) {
	rows, err := config.DB.Query(
		`SELECT id, nombre, cantidad, fecha_vencimiento, categoria_id, usuario_id
		 FROM productos
		 WHERE usuario_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []models.Producto

	for rows.Next() {
		var p models.Producto
		err := rows.Scan(
			&p.ID,
			&p.Nombre,
			&p.Cantidad,
			&p.FechaVencimiento,
			&p.CategoriaID,
			&p.UsuarioID,
		)
		if err != nil {
			return nil, err
		}
		productos = append(productos, p)
	}

	return productos, nil
}

func (r *ProductoRepository) GetByID(id, userID int) (*models.Producto, error) {
	var p models.Producto

	err := config.DB.QueryRow(`
		SELECT id, nombre, cantidad, fecha_vencimiento, categoria_id, usuario_id
		FROM productos
		WHERE id = ? AND usuario_id = ?`,
		id, userID,
	).Scan(
		&p.ID,
		&p.Nombre,
		&p.Cantidad,
		&p.FechaVencimiento,
		&p.CategoriaID,
		&p.UsuarioID,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *ProductoRepository) Update(id int, p *models.Producto, userID int) error {
	_, err := config.DB.Exec(`
		UPDATE productos
		SET nombre = ?, cantidad = ?, fecha_vencimiento = ?, categoria_id = ?
		WHERE id = ? AND usuario_id = ?`,
		p.Nombre,
		p.Cantidad,
		p.FechaVencimiento,
		p.CategoriaID,
		id,
		userID,
	)

	return err
}

func (r *ProductoRepository) Delete(id, userID int) error {
	_, err := config.DB.Exec(`
		DELETE FROM productos
		WHERE id = ? AND usuario_id = ?`,
		id,
		userID,
	)

	return err
}
