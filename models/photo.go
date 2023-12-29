package models


import "gorm.io/gorm"

type Photo struct {
	gorm.Model

	ID			uint	`gorm:"primaryKey,not null"`
	Title		string	`gorm:"not null"`
	Caption		string	
	PhotoURL	string	`gorm:"not null"`
	UserID		uint	`gorm:"not null"`
	User 		User	`gorm:"foreignKey:UserID`
}