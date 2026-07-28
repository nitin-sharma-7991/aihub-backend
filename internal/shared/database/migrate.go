package database

import (
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/user/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	models := []any{
		&model.User{},
	}

	return db.AutoMigrate(models...)
}
