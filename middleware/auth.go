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
		// Mendapatkan token dari header Authorization
		tokenString := helpers.ExtractToken(c.Request)

		// Verifikasi token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Menggunakan secret key untuk verifikasi
			return []byte("mpizesterisjdjksjdskdjansakj123"), nil
		})

		if err != nil || !token.Valid {
			// Token tidak valid atau terjadi kesalahan
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Token valid, lanjutkan ke handler selanjutnya
		c.Next()
	}
}