package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// User represents a user in the appointment system
type User struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null;size:255" json:"name" validate:"required,min=2,max=255"`
	Email       string         `gorm:"uniqueIndex;not null;size:255" json:"email" validate:"required,email,max=255"`
	PhoneNumber string         `gorm:"size:20" json:"phone_number" validate:"omitempty,max=20"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "users"
}

// Validate validates the user struct using the validator package
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// BeforeCreate is a GORM hook that runs before creating a user
func (u *User) BeforeCreate(tx *gorm.DB) error {
	return u.Validate()
}

// BeforeUpdate is a GORM hook that runs before updating a user
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	return u.Validate()
}
