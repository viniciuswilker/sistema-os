package routes

import (
	"sistema-os/internal/handlers"
	"sistema-os/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigurarRotas(r *gin.Engine, db *gorm.DB) {

	clienteRepo := repository.NovoClienteRepository(db)
	clienteHandler := handlers.NovoClienteHandler(clienteRepo)

	api := r.Group("/api/v1")
	{
		clientes := api.Group("/clientes")
		{
			clientes.POST("/", clienteHandler.CriarCliente)
			clientes.GET("/", clienteHandler.ListarClientes)
			clientes.GET("/:id", clienteHandler.BuscarCliente)

		}

	}

}
