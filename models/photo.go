package models
import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	ID			uint	`gorm:"primaryKey,not null"`
	Title		string	`gorm:"not null"`
	Caption		string	
	PhotoURL 	string 	`gorm:"not null" json:"photo_url"`
	UserID		uint	`gorm:"not null" json:"user_id"`
	User 		User	`gorm:"foreignKey:UserID"`
}