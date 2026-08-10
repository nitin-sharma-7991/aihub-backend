package repository

import (
	"context"

	rolePermissionModel "github.com/nitin-sharma-7991/aihub-backend/internal/modules/role_permission/model"
	"gorm.io/gorm"
)

type RolePermissionRepository interface {
	Exists(
		ctx context.Context,
		roleID uint,
		permissionID uint,
	) (bool, error)

	Create(
		ctx context.Context,
		rolePermission *rolePermissionModel.RolePermission,
	) error

	Delete(
		ctx context.Context,
		roleID uint,
		permissionID uint,
	) error
}

type rolePermissionRepository struct {
	db *gorm.DB
}

func NewRolePermissionRepository(
	db *gorm.DB,
) RolePermissionRepository {
	return &rolePermissionRepository{
		db: db,
	}
}

// Check whether a permission is assigned to a role.
func (r *rolePermissionRepository) Exists(
	ctx context.Context,
	roleID uint,
	permissionID uint,
) (bool, error) {

	var count int64

	err := r.db.
		WithContext(ctx).
		Model(&rolePermissionModel.RolePermission{}).
		Where(
			"role_id = ? AND permission_id = ?",
			roleID,
			permissionID,
		).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Assign permission to role.
func (r *rolePermissionRepository) Create(
	ctx context.Context,
	rolePermission *rolePermissionModel.RolePermission,
) error {

	return r.db.
		WithContext(ctx).
		Create(rolePermission).
		Error
}

// Remove permission from role.
func (r *rolePermissionRepository) Delete(
	ctx context.Context,
	roleID uint,
	permissionID uint,
) error {

	return r.db.
		WithContext(ctx).
		Where(
			"role_id = ? AND permission_id = ?",
			roleID,
			permissionID,
		).
		Delete(&rolePermissionModel.RolePermission{}).
		Error
}
