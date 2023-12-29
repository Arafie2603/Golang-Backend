// controllers/user_controller.go

package controllers

import (
    "github.com/gin-gonic/gin"
    "finpro-golang2/models"
    "finpro-golang2/helpers"
    "net/http"
    "strconv"
	"fmt"
    "finpro-golang2/database"
)

func Register(c *gin.Context) {
    // Baca data yang diterima dari body request
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash password sebelum menyimpan ke database
    hashedPassword, err := helpers.HashPassword(user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }
    user.Password = hashedPassword

    // Simpan user ke database
    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    // Generate token setelah registrasi berhasil
    token, err := helpers.GenerateToken(int(user.ID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    // Kirim respons dengan token
    c.JSON(http.StatusOK, gin.H{"token": token})
}

// controllers/user_controller.go

// ...

func Login(c *gin.Context) {
    // Baca data yang diterima dari body request
    var userInput models.User
    if err := c.ShouldBindJSON(&userInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Cari pengguna berdasarkan email
    var userFromDB models.User
    if err := database.DB.Where("email = ?", userInput.Email).First(&userFromDB).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Verifikasi password
    if !helpers.CheckPasswordHash(userInput.Password, userFromDB.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
        return
    }

    // Generate token setelah login berhasil
    token, err := helpers.GenerateToken(int(userFromDB.ID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    // Kirim respons dengan token
    c.JSON(http.StatusOK, gin.H{"token": token})
}


// ...



func UpdateUser(c *gin.Context) {
	// Mendapatkan ID pengguna dari parameter URL
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Mendapatkan pengguna dari database berdasarkan ID
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Baca data yang diterima dari body request
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan perubahan ke database
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Kirim respons
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}


	if err := database.DB.Where("id = ?", userID).Delete(&models.User{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		fmt.Println("Error deleting user:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

	fmt.Println("User deleted successfully")
}


