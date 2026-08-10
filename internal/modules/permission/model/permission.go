package model

import "time"

type Permission struct {
	ID uint `gorm:"primaryKey"`

	Name        string `gorm:"size:100;not null;uniqueIndex"`
	Description string `gorm:"size:255"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
