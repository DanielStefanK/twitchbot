package storage

import (
	"gorm.io/gorm"
)

// Migrate every model to the new DB
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&ChatConfig{}, &CustomCommand{})
}
