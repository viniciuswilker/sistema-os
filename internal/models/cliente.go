package models

import (
	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model
	Nome     string `gorm:"not null" json:"nome"`
	Email    string `gorm:"unique" json:"email"`
	Telefone string `json:"telefone"`
	Endereco string `json:"endereco"`
}

