package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Nome      string `gorm:"not null" json:"nome"`
	Sobrenome string `gorm:"not null" json:"sobrenome"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `gorm:"not null" json:"-"`
	RG        string `gorm:"unique;not null" json:"rg"`
	Endereco  string `json:"endereco"`
}

func (u *Usuario) BeforeSave(tx *gorm.DB) (err error) {

	if len(u.Password) > 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hash)
	}
	return nil
}
