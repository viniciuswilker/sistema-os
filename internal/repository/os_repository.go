package repository

import (
	"sistema-os/internal/models"

	"gorm.io/gorm"
)

type OSRepository struct {
	db *gorm.DB
}

func NovoOSRepository(db *gorm.DB) *OSRepository {
	return &OSRepository{db: db}

}

func (r *OSRepository) Criar(os models.OrdemServico) error {
	return r.db.Create(os).Error
}

func (r *OSRepository) Listar() ([]models.OrdemServico, error) {
	var listaOS []models.OrdemServico

	err := r.db.Preload("Cliente").Find(&listaOS).Error
	return listaOS, err

}

func (r *OSRepository) BuscarPorID(id uint) (*models.OrdemServico, error) {
	var os models.OrdemServico
	err := r.db.Preload("Cliente").First(&os, id).Error
	return &os, err
}

func (r *OSRepository) AtualizarStatus(id uint, status string, valor float64) error {
	return r.db.Model(&models.OrdemServico{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status": status,
		"valor":  valor,
	}).Error
}
