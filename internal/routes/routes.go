package routes

import (
	"net/http"
	"sistema-os/internal/handlers"
	"sistema-os/internal/middleware"
	"sistema-os/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigurarRotas(r *gin.Engine, db *gorm.DB) {

	r.Static("/assets", "./assets")
	r.Static("/uploads", "./uploads")

	clienteRepo := repository.NovoClienteRepository(db)
	osRepo := repository.NovoOSRepository(db)

	clienteHandler := handlers.NovoClienteHandler(clienteRepo)
	osHandler := handlers.NovoOSHandler(osRepo)
	authHandler := handlers.NovoAuthHandler(db)

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{"titulo": "Login"})
	})

	r.GET("/dashboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", gin.H{"titulo": "Painel de Controle"})
	})

	api := r.Group("/api/v1")
	{
		// Públicas
		api.POST("/login", authHandler.Login)
		api.POST("/registrar", authHandler.Registrar)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			// Grupo Clientes
			clientes := protected.Group("/clientes")
			{
				clientes.POST("/", clienteHandler.CriarCliente)
				clientes.GET("/", clienteHandler.ListarClientes)
				clientes.GET("/:id", clienteHandler.BuscarCliente)
			}

			// Grupo Ordens de Serviço
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
