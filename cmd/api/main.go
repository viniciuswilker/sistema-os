package main

import (
	"sistema-os/internal/database"
	"sistema-os/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {

	database.InitDB()

	database.DB.AutoMigrate(&models.Cliente{}, &models.Usuario{})

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong - Sistema de OS Online",
		})
	})

	r.Run(":8080")

}
