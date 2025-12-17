package models

import "gorm.io/gorm"

type OrdemServico struct {
	gorm.Model

	ClientID uint    `gorm:"not null" json:"cliente_id"`
	Cliente  Cliente `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"cliente,omitempty"`

	UsuarioID uint    `gorm:"not null" json:"usuario_id"`
	Usuario   Usuario `gorm:"usuario, omitempty"`

	Aparelho    string  `gorm:"not null" json:"aparelho"`
	Defeito     string  `gorm:"not null" json:"defeito"`
	Status      string  `gorm:"default:'Recebido'" json:"status"`
	Valor       float64 `json:"valor"`
	Observacoes string  `json:"observacoes"`
	Fotos       []Foto  `gorm:"foreignKey:OrdemServicoID" json:"fotos"`
}
