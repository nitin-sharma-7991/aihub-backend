package model

import "time"

type Organization struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:100;not null"`
	Slug        string `gorm:"size:100;uniqueIndex;not null"`
	Description string
	CreatedBy   uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
