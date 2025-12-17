package main

import (
	"sistema-os/internal/database"
	"sistema-os/internal/models"
	"sistema-os/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.InitDB()

	database.DB.AutoMigrate(&models.Cliente{}, &models.Usuario{}, &models.OrdemServico{}, &models.Cliente{})

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	routes.ConfigurarRotas(r, database.DB)

	r.Run(":8080")

}
