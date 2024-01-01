// middlewares/auth.go

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"finpro-golang2/helpers"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := helpers.ExtractToken(c.Request)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("mpizesterisjdjksjdskdjansakj123"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}