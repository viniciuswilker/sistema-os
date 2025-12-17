package handlers

import (
	"net/http"
	"sistema-os/internal/auth"
	"sistema-os/internal/models"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	db *gorm.DB
}

func NovoAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var credenciais struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credenciais); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	var usuario models.Usuario
	if err := h.db.Where("email = ?", credenciais.Email).First(&usuario).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "Credenciais inválidas"})
		return
	}

	if !auth.CheckPassword(credenciais.Password, usuario.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "Credenciais inválidas"})
		return
	}

	token, err := auth.GenerateToken(usuario.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) Registrar(c *gin.Context) {
	var usuario models.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := h.db.Create(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao criar usuário. Email ou RG já existem?"})
		return
	}

	usuario.Password = ""
	c.JSON(http.StatusCreated, usuario)
}
