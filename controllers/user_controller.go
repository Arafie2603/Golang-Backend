package controllers

import (
    "github.com/gin-gonic/gin"
    "finpro-golang2/models"
    "finpro-golang2/helpers"
    "net/http"
    "strconv"
	"fmt"
    "finpro-golang2/database"
	"log"
)

func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash password
    hashedPassword, err := helpers.HashPassword(user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }
    user.Password = hashedPassword

    // Buat pengguna baru
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

    // Kirim respons dengan informasi berhasil
    c.JSON(http.StatusOK, gin.H{"token": token, "message": "Registration successful"})
}


func Login(c *gin.Context) {
	var userInput models.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFromDB models.User
	if err := database.DB.Where("email = ?", userInput.Email).First(&userFromDB).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if !helpers.CheckPasswordHash(userInput.Password, userFromDB.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	token, err := helpers.GenerateToken(int(userFromDB.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	log.Printf("User with ID %d has logged in", userFromDB.ID)

	c.JSON(http.StatusOK, gin.H{"token": token, "message":"Login succesful"})
}


func UpdateUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Bind data dari request ke struct user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Jika ada perubahan pada password, hash password baru
	if len(user.Password) > 0 {
		hashedPassword, err := helpers.HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
			return
		}
		user.Password = hashedPassword
	}

	// Simpan perubahan ke basis data
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

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

