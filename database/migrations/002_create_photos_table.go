// database/migrations/002_create_photos_table.go

package migrations

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Caption  string
	PhotoURL string `gorm:"not null"`
	UserID   uint   `gorm:"not null"`
}

func MigratePhotos(db *gorm.DB) {
	db.AutoMigrate(&Photo{})
}
