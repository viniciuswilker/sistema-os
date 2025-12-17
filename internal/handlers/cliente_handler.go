package handlers

import (
	"net/http"
	"sistema-os/internal/models"
	"sistema-os/internal/repository"
	"strconv"

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
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao criar cliente"})
		return
	}

	c.JSON(http.StatusCreated, cliente)
}

func (h *ClienteHandler) ListarClientes(c *gin.Context) {
	clientes, err := h.repo.Listar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao listar clientes"})
		return
	}
	c.JSON(http.StatusOK, clientes)
}

func (h *ClienteHandler) BuscarCliente(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	cliente, err := h.repo.BuscarPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Cliente não encontrado"})
		return
	}

	c.JSON(http.StatusOK, cliente)
}
