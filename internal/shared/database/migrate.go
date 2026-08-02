package database

import (
	membershipModule "github.com/nitin-sharma-7991/aihub-backend/internal/modules/membership/model"
	orgModule "github.com/nitin-sharma-7991/aihub-backend/internal/modules/organization/model"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/user/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	models := []any{
		&model.User{},
		&orgModule.Organization{},
		&membershipModule.Membership{},
	}

	return db.AutoMigrate(models...)
}
