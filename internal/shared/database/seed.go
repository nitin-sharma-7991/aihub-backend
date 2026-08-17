package database

import (
	"errors"

	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/permission/model"
	roleModel "github.com/nitin-sharma-7991/aihub-backend/internal/modules/role/model"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {

	if err := seedRoles(db); err != nil {
		return err
	}

	if err := seedPermissions(db); err != nil {
		return err
	}

	if err := seedRolePermissions(db); err != nil {
		return err
	}

	return nil
}

func seedRoles(db *gorm.DB) error {

	roles := []roleModel.Role{
		{
			Name:        "super_admin",
			Description: "System administrator with full access",
		},
		{
			Name:        "organization_owner",
			Description: "Owner of an organization",
		},
		{
			Name:        "admin",
			Description: "Organization administrator",
		},
		{
			Name:        "developer",
			Description: "Developer with limited organization access",
		},
		{
			Name:        "viewer",
			Description: "Read-only access",
		},
	}

	for _, role := range roles {

		var existing roleModel.Role

		err := db.
			Where("name = ?", role.Name).
			First(&existing).
			Error

		if err == nil {
			continue
		}

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err := db.Create(&role).Error; err != nil {
			return err
		}
	}

	return nil
}

func seedPermissions(db *gorm.DB) error {

	permissions := []model.Permission{

		// Organization
		{
			Name:        "organization:create",
			Description: "Create organization",
		},
		{
			Name:        "organization:read",
			Description: "Read organization",
		},
		{
			Name:        "organization:update",
			Description: "Update organization",
		},
		{
			Name:        "organization:delete",
			Description: "Delete organization",
		},

		// Membership
		{
			Name:        "membership:create",
			Description: "Create organization membership",
		},
		{
			Name:        "membership:read",
			Description: "Read organization memberships",
		},
		{
			Name:        "membership:update",
			Description: "Update organization membership",
		},
		{
			Name:        "membership:delete",
			Description: "Delete organization membership",
		},

		// Roles
		{
			Name:        "role:create",
			Description: "Create role",
		},
		{
			Name:        "role:read",
			Description: "Read roles",
		},
		{
			Name:        "role:update",
			Description: "Update role",
		},
		{
			Name:        "role:delete",
			Description: "Delete role",
		},
	}

	for _, permission := range permissions {

		var existing model.Permission

		err := db.
			Where("name = ?", permission.Name).
			First(&existing).
			Error

		if err == nil {
			continue
		}

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err := db.Create(&permission).Error; err != nil {
			return err
		}
	}

	return nil
}

func seedRolePermissions(db *gorm.DB) error {

	// --------------------------------------------------
	// Super Admin
	// --------------------------------------------------

	superAdminPermissions := []string{
		"organization:create",
		"organization:read",
		"organization:update",
		"organization:delete",

		"membership:create",
		"membership:read",
		"membership:update",
		"membership:delete",

		"role:create",
		"role:read",
		"role:update",
		"role:delete",
	}

	if err := assignPermissions(
		db,
		"super_admin",
		superAdminPermissions,
	); err != nil {
		return err
	}

	// --------------------------------------------------
	// Organization Owner
	// --------------------------------------------------

	ownerPermissions := []string{
		"organization:read",
		"organization:update",

		"membership:create",
		"membership:read",
		"membership:update",
		"membership:delete",
	}

	if err := assignPermissions(
		db,
		"organization_owner",
		ownerPermissions,
	); err != nil {
		return err
	}

	// --------------------------------------------------
	// Admin
	// --------------------------------------------------

	adminPermissions := []string{
		"organization:read",

		"membership:create",
		"membership:read",
		"membership:update",
		"membership:delete",
	}

	if err := assignPermissions(
		db,
		"admin",
		adminPermissions,
	); err != nil {
		return err
	}

	// --------------------------------------------------
	// Developer
	// --------------------------------------------------

	developerPermissions := []string{
		"organization:read",

		"membership:read",
	}

	if err := assignPermissions(
		db,
		"developer",
		developerPermissions,
	); err != nil {
		return err
	}

	// --------------------------------------------------
	// Viewer
	// --------------------------------------------------

	viewerPermissions := []string{
		"organization:read",
	}

	if err := assignPermissions(
		db,
		"viewer",
		viewerPermissions,
	); err != nil {
		return err
	}

	return nil
}

func assignPermissions(
	db *gorm.DB,
	roleName string,
	permissionNames []string,
) error {

	var role roleModel.Role

	if err := db.
		Where("name = ?", roleName).
		First(&role).
		Error; err != nil {
		return err
	}

	for _, permissionName := range permissionNames {

		var permission model.Permission

		if err := db.
			Where("name = ?", permissionName).
			First(&permission).
			Error; err != nil {
			return err
		}

		var count int64

		err := db.
			Table("role_permissions").
			Where(
				"role_id = ? AND permission_id = ?",
				role.ID,
				permission.ID,
			).
			Count(&count).
			Error

		if err != nil {
			return err
		}

		if count > 0 {
			continue
		}

		if err := db.
			Table("role_permissions").
			Create(map[string]interface{}{
				"role_id":       role.ID,
				"permission_id": permission.ID,
			}).
			Error; err != nil {
			return err
		}
	}

	return nil
}
