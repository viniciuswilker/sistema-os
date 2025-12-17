package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("sistema.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados:", err)

	}

	DB = DB.Debug()

}
