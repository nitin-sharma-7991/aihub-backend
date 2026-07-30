package model

import "time"

type User struct {
	ID uint `gorm:"primaryKey"`
	// Temporary until Dynamic RBAC is implemented
	Role string `gorm:"size:50;not null;default:'viewer'"`

	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:150;uniqueIndex;not null"`
	Password  string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
