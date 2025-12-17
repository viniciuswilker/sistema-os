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

type Usuario struct {
	gorm.Model
	Nome      string `gorm:"not null" json:"nome"`
	Sobrenome string `gorm:"not null" json:"sobrenome"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `gorm:"not null" json:"-"`
	RG        string `gorm:"unique;not null" json:"rg"`
	Endereco  string `json:"endereco"`
}
