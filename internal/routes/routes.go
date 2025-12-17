package routes

import (
	"sistema-os/internal/handlers"
	"sistema-os/internal/middleware" // <--- Importante: Importar o middleware
	"sistema-os/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigurarRotas(r *gin.Engine, db *gorm.DB) {

	// REPOS
	clienteRepo := repository.NovoClienteRepository(db)
	osRepo := repository.NovoOSRepository(db)

	// HANDLERS
	clienteHandler := handlers.NovoClienteHandler(clienteRepo)
	osHandler := handlers.NovoOSHandler(osRepo)
	authHandler := handlers.NovoAuthHandler(db)

	r.Static("/assets", "./uploads")

	api := r.Group("/api/v1")
	{

		api.POST("/login", authHandler.Login)
		api.POST("/registrar", authHandler.Registrar)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			clientes := protected.Group("/clientes")
			{
				clientes.POST("/", clienteHandler.CriarCliente)
				clientes.GET("/", clienteHandler.ListarClientes)
				clientes.GET("/:id", clienteHandler.BuscarCliente)
			}

			oss := protected.Group("/os")
			{
				oss.POST("/", osHandler.CriarOS)
				oss.GET("/", osHandler.ListarOS)
				oss.GET("/:id", osHandler.BuscarOS)
				oss.POST("/:id/fotos", osHandler.UploadFoto)
			}
		}
	}
}
