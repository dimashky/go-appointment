package models

import (
	"gorm.io/gorm"
)

// AllModels returns a slice of all models for migration purposes
func AllModels() []interface{} {
	return []interface{}{
		&User{},
	}
}

// AutoMigrate runs automatic migration for all models
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(AllModels()...)
}
