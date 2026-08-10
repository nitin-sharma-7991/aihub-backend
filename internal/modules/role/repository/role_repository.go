package repository

import (
	"context"

	roleModel "github.com/nitin-sharma-7991/aihub-backend/internal/modules/role/model"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/pagination"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(
		ctx context.Context,
		role *roleModel.Role,
	) error

	FindAll(
		ctx context.Context,
		req pagination.Request,
	) ([]roleModel.Role, int64, error)

	FindByID(
		ctx context.Context,
		id uint,
	) (*roleModel.Role, error)

	FindByName(
		ctx context.Context,
		name string,
	) (*roleModel.Role, error)

	Update(
		ctx context.Context,
		role *roleModel.Role,
	) error

	Delete(
		ctx context.Context,
		id uint,
	) error
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(
	db *gorm.DB,
) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

// Create Role
func (r *roleRepository) Create(
	ctx context.Context,
	role *roleModel.Role,
) error {
	return r.db.
		WithContext(ctx).
		Create(role).
		Error
}

// Find All Roles
func (r *roleRepository) FindAll(
	ctx context.Context,
	req pagination.Request,
) ([]roleModel.Role, int64, error) {

	var (
		roles []roleModel.Role
		total int64
	)

	db := r.db.WithContext(ctx)

	if err := db.
		Model(&roleModel.Role{}).
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := db.
		Order("id ASC").
		Limit(req.Limit).
		Offset(req.Offset()).
		Find(&roles).
		Error; err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}

// Find Role By ID
func (r *roleRepository) FindByID(
	ctx context.Context,
	id uint,
) (*roleModel.Role, error) {

	var role roleModel.Role

	err := r.db.
		WithContext(ctx).
		First(&role, id).
		Error

	if err != nil {
		return nil, err
	}

	return &role, nil
}

// Find Role By Name
func (r *roleRepository) FindByName(
	ctx context.Context,
	name string,
) (*roleModel.Role, error) {

	var role roleModel.Role

	err := r.db.
		WithContext(ctx).
		Where("name = ?", name).
		First(&role).
		Error

	if err != nil {
		return nil, err
	}

	return &role, nil
}

// Update Role
func (r *roleRepository) Update(
	ctx context.Context,
	role *roleModel.Role,
) error {
	return r.db.
		WithContext(ctx).
		Save(role).
		Error
}

// Delete Role
func (r *roleRepository) Delete(
	ctx context.Context,
	id uint,
) error {
	return r.db.
		WithContext(ctx).
		Delete(&roleModel.Role{}, id).
		Error
}
