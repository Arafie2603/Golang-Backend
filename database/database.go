package database

import (
	"gorm.io/driver/mysql"
	"finpro-golang2/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/latihan_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	DB = db
	DB.AutoMigrate(&models.User{}, &models.Photo{})
}