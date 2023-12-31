package helpers

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"errors"
)

var secretKey = []byte("mpizesterisjdjksjdskdjansakj123")

func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}

	// Membuat token dengan klaim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Menandatangani token dengan secret key dan mendapatkan string token
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractUserIDFromToken(tokenString string) (int, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(secretKey), nil
    })

    if err != nil {
        return 0, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        if userID, exists := claims["user_id"].(float64); exists {
            return int(userID), nil
        }
    }

    return 0, errors.New("Invalid token or user ID not found in claims")
}

// VerifyToken memeriksa keaslian token JWT dan mengembalikan klaim token jika valid
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	// Mendeklarasikan fungsi pembacaan klaim token
	claims := jwt.MapClaims{}

	// Verifikasi token
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}