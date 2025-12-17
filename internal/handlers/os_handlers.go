package handlers

import (
	"net/http"
	"sistema-os/internal/models"
	"sistema-os/internal/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OSHandler struct {
	repo *repository.OSRepository
}

func NovoOSHandler(repo *repository.OSRepository) *OSHandler {
	return &OSHandler{repo: repo}
}

func (h *OSHandler) CriarOS(c *gin.Context) {
	var os models.OrdemServico
	if err := c.ShouldBindJSON(&os); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := h.repo.Criar(&os); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao criar OS"})
		return
	}

	c.JSON(http.StatusCreated, os)
}

func (h *OSHandler) ListarOS(c *gin.Context) {
	lista, err := h.repo.Listar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar ordens"})
		return
	}
	c.JSON(http.StatusOK, lista)
}

func (h *OSHandler) BuscarOS(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	os, err := h.repo.BuscarPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "OS não encontrada"})
		return
	}
	c.JSON(http.StatusOK, os)
}
