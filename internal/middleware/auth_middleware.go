package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("sua_chave_secreta_super_dificil")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearerSchema = "Bearer "
		header := c.GetHeader("Authorization")

		if header == "" {
			c.AbortWithStatusJSON(401, gin.H{"erro": "Token não fornecido"})
			return
		}

		tokenString := header[len(bearerSchema):]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"erro": "Token inválido ou expirado"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("usuario_id", claims["sub"])
			c.Set("usuario_tipo", claims["tipo"])
		}

		c.Next()
	}
}
