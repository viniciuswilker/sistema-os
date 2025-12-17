package handlers

import (
	"net/http"
	"sistema-os/internal/models"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	db *gorm.DB
}

func NovoAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

// /api/v1/login
var jwtKey = []byte("sua_chave_secreta_super_dificil")

func (h *AuthHandler) Login(c *gin.Context) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(400, gin.H{"erro": "Dados inválidos"})
		return
	}

	var user models.Usuario
	if err := h.db.Where("email = ?", creds.Email).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"erro": "Usuário ou senha incorretos"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		c.JSON(400, gin.H{"erro": "Usuário ou senha incorretos"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.MapClaims{
		"sub": user.ID,
		"exp": expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(500, gin.H{"erro": "Erro ao gerar token"})
		return
	}

	c.JSON(200, gin.H{
		"token": tokenString,
	})
}

// /api/v1/registrar
func (h *AuthHandler) Registrar(c *gin.Context) {
	var req struct {
		Nome     string `json:"Nome" binding:"required"`
		Email    string `json:"Email" binding:"required,email"`
		Password string `json:"Password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao processar senha"})
		return
	}

	novoUsuario := models.Usuario{
		Nome:     req.Nome,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := h.db.Create(&novoUsuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao criar usuário"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":   novoUsuario.ID,
		"nome": novoUsuario.Nome,
		"msg":  "Usuário criado com sucesso",
	})
}
