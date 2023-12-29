// helpers/bcrypt.go
package helpers

import "golang.org/x/crypto/bcrypt"

// HashPassword digunakan untuk menghash password menggunakan bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(inputPassword, hashedPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
    return err == nil
}
