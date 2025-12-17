package repository

import (
	"sistema-os/internal/models"

	"gorm.io/gorm"
)

type ClienteRepository struct {
	db *gorm.DB
}

func NovoClienteRepository(db *gorm.DB) *ClienteRepository {
	return &ClienteRepository{db: db}
}

func (r *ClienteRepository) Criar(cliente *models.Cliente) error {
	return r.db.Create(cliente).Error
}

func (r *ClienteRepository) Listar() ([]models.Cliente, error) {
	var clientes []models.Cliente
	err := r.db.Find(&clientes).Error
	return clientes, err
}

func (r *ClienteRepository) BuscarPorID(id uint) (*models.Cliente, error) {
	var cliente models.Cliente
	err := r.db.First(&cliente, id).Error
	return &cliente, err
}

func (r *ClienteRepository) Atualizar(cliente *models.Cliente) error {
	return r.db.Save(cliente).Error
}
