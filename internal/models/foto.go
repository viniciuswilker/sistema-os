package models

import "gorm.io/gorm"

type Foto struct {
	gorm.Model
	OrdemServicoID uint   `json:"os_id"`
	Caminho        string `json:"caminho"`
}
