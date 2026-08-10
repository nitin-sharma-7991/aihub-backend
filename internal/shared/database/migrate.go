package database

import (
	membershipModule "github.com/nitin-sharma-7991/aihub-backend/internal/modules/membership/model"
	orgModule "github.com/nitin-sharma-7991/aihub-backend/internal/modules/organization/model"
	permissionModule "github.com/nitin-sharma-7991/aihub-backend/internal/modules/permission/model"
	roleModule "github.com/nitin-sharma-7991/aihub-backend/internal/modules/role/model"
	rolePermissionModule "github.com/nitin-sharma-7991/aihub-backend/internal/modules/role_permission/model"
	userModule "github.com/nitin-sharma-7991/aihub-backend/internal/modules/user/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	models := []any{
		&userModule.User{},
		&orgModule.Organization{},
		&roleModule.Role{},
		&permissionModule.Permission{},
		&rolePermissionModule.RolePermission{},
		&membershipModule.Membership{},
	}

	return db.AutoMigrate(models...)
}
