package handlers

import (
	"net/http"
	"sistema-os/internal/models"
	"sistema-os/internal/repository"

	"github.com/gin-gonic/gin"
)

type ClienteHandler struct {
	repo *repository.ClienteRepository
}

func NovoClienteHandler(repo *repository.ClienteRepository) *ClienteHandler {
	return &ClienteHandler{repo: repo}
}

func (h *ClienteHandler) CriarCliente(c *gin.Context) {
	var cliente models.Cliente

	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := h.repo.Criar(&cliente); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao criar o cliente"})
		return
	}

	c.JSON(http.StatusCreated, cliente)

}
