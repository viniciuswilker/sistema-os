package routes

import (
	"sistema-os/internal/handlers"
	"sistema-os/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigurarRotas(r *gin.Engine, db *gorm.DB) {

	// repos
	clienteRepo := repository.NovoClienteRepository(db)
	osRepo := repository.NovoOSRepository(db)

	// handlers
	clienteHandler := handlers.NovoClienteHandler(clienteRepo)
	osHandler := handlers.NovoOSHandler(osRepo)

	api := r.Group("/api/v1")
	{
		clientes := api.Group("/clientes")
		{
			clientes.POST("/", clienteHandler.CriarCliente)
			clientes.GET("/", clienteHandler.ListarClientes)
			clientes.GET("/:id", clienteHandler.BuscarCliente)

		}

		oss := api.Group("/os")

		{
			oss.POST("/", osHandler.CriarOS)
			oss.GET("/", osHandler.ListarOS)
			oss.GET("/:id", osHandler.BuscarOS)
		}

	}

}
