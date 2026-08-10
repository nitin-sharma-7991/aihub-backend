package repository

import (
	"context"

	permissionModel "github.com/nitin-sharma-7991/aihub-backend/internal/modules/permission/model"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/pagination"

	"gorm.io/gorm"
)

type PermissionRepository interface {
	Create(
		ctx context.Context,
		permission *permissionModel.Permission,
	) error

	FindAll(
		ctx context.Context,
		req pagination.Request,
	) ([]permissionModel.Permission, int64, error)

	FindByID(
		ctx context.Context,
		id uint,
	) (*permissionModel.Permission, error)

	FindByName(
		ctx context.Context,
		name string,
	) (*permissionModel.Permission, error)

	Update(
		ctx context.Context,
		permission *permissionModel.Permission,
	) error

	Delete(
		ctx context.Context,
		id uint,
	) error

	FindByRoleID(
		ctx context.Context,
		roleID uint,
	) ([]permissionModel.Permission, error)
}

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(
	db *gorm.DB,
) PermissionRepository {
	return &permissionRepository{
		db: db,
	}
}

// Create Permission
func (r *permissionRepository) Create(
	ctx context.Context,
	permission *permissionModel.Permission,
) error {
	return r.db.
		WithContext(ctx).
		Create(permission).
		Error
}

// Find All Permissions
func (r *permissionRepository) FindAll(
	ctx context.Context,
	req pagination.Request,
) ([]permissionModel.Permission, int64, error) {

	var (
		permissions []permissionModel.Permission
		total       int64
	)

	db := r.db.WithContext(ctx)

	if err := db.
		Model(&permissionModel.Permission{}).
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := db.
		Order("id ASC").
		Limit(req.Limit).
		Offset(req.Offset()).
		Find(&permissions).
		Error; err != nil {
		return nil, 0, err
	}

	return permissions, total, nil
}

// Find Permission By ID
func (r *permissionRepository) FindByID(
	ctx context.Context,
	id uint,
) (*permissionModel.Permission, error) {

	var permission permissionModel.Permission

	err := r.db.
		WithContext(ctx).
		First(&permission, id).
		Error

	if err != nil {
		return nil, err
	}

	return &permission, nil
}

// Find Permission By Name
func (r *permissionRepository) FindByName(
	ctx context.Context,
	name string,
) (*permissionModel.Permission, error) {

	var permission permissionModel.Permission

	err := r.db.
		WithContext(ctx).
		Where("name = ?", name).
		First(&permission).
		Error

	if err != nil {
		return nil, err
	}

	return &permission, nil
}

// Update Permission
func (r *permissionRepository) Update(
	ctx context.Context,
	permission *permissionModel.Permission,
) error {
	return r.db.
		WithContext(ctx).
		Save(permission).
		Error
}

// Delete Permission
func (r *permissionRepository) Delete(
	ctx context.Context,
	id uint,
) error {
	return r.db.
		WithContext(ctx).
		Delete(&permissionModel.Permission{}, id).
		Error
}

// Find Permissions Assigned To Role
func (r *permissionRepository) FindByRoleID(
	ctx context.Context,
	roleID uint,
) ([]permissionModel.Permission, error) {

	var permissions []permissionModel.Permission

	err := r.db.
		WithContext(ctx).
		Table("permissions").
		Joins(
			"JOIN role_permissions ON role_permissions.permission_id = permissions.id",
		).
		Where(
			"role_permissions.role_id = ?",
			roleID,
		).
		Order("permissions.id ASC").
		Find(&permissions).
		Error

	if err != nil {
		return nil, err
	}

	return permissions, nil
}
