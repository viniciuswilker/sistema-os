package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sistema-os/internal/models"
	"sistema-os/internal/repository"
	"strconv"
	"time"

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

	usuarioID, exists := c.Get("usuarioID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "Usuário não autenticado"})
		return
	}

	os.UsuarioID = usuarioID.(uint)

	if os.UsuarioID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID de usuário inválido"})
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

func (h *OSHandler) UploadFoto(c *gin.Context) {
	idParam := c.Param("id")
	osID, _ := strconv.Atoi(idParam)

	file, err := c.FormFile("foto")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Arquivo não enviado"})
		return
	}

	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", os.ModePerm)
	}

	ext := filepath.Ext(file.Filename)
	novoNome := fmt.Sprintf("%d_%d%s", osID, time.Now().Unix(), ext)
	caminhoDestino := "uploads/" + novoNome

	if err := c.SaveUploadedFile(file, caminhoDestino); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao salvar arquivo"})
		return
	}

	foto := models.Foto{
		OrdemServicoID: uint(osID),
		Caminho:        caminhoDestino,
	}

	if err := h.repo.AdicionarFoto(&foto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao registrar foto no banco"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Upload realizado", "foto": foto})
}
