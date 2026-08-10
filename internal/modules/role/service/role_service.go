package service

import (
	"context"
	"errors"
	"strings"

	roleModel "github.com/nitin-sharma-7991/aihub-backend/internal/modules/role/model"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/role/repository"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/apperrors"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/pagination"

	"gorm.io/gorm"
)

type RoleService interface {
	Create(
		ctx context.Context,
		name string,
		description string,
	) (*roleModel.Role, error)

	FindAll(
		ctx context.Context,
		req pagination.Request,
	) ([]roleModel.Role, pagination.Meta, error)

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
		id uint,
		name string,
		description string,
	) (*roleModel.Role, error)

	Delete(
		ctx context.Context,
		id uint,
	) error
}

type roleService struct {
	roleRepo repository.RoleRepository
}

func NewRoleService(
	roleRepo repository.RoleRepository,
) RoleService {
	return &roleService{
		roleRepo: roleRepo,
	}
}

// Create Role
func (s *roleService) Create(
	ctx context.Context,
	name string,
	description string,
) (*roleModel.Role, error) {

	name = strings.TrimSpace(name)
	description = strings.TrimSpace(description)

	if name == "" {
		return nil, apperrors.ErrInvalidRoleName
	}

	// Prevent duplicate role names.
	existing, err := s.roleRepo.FindByName(ctx, name)

	if err == nil && existing != nil {
		return nil, apperrors.ErrRoleAlreadyExists
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	role := &roleModel.Role{
		Name:        name,
		Description: description,
	}

	if err := s.roleRepo.Create(ctx, role); err != nil {
		return nil, err
	}

	return role, nil
}

// Find All Roles
func (s *roleService) FindAll(
	ctx context.Context,
	req pagination.Request,
) ([]roleModel.Role, pagination.Meta, error) {

	req.Normalize()

	roles, total, err := s.roleRepo.FindAll(
		ctx,
		req,
	)

	if err != nil {
		return nil, pagination.Meta{}, err
	}

	meta := pagination.NewMeta(req, total)

	return roles, meta, nil
}

// Find Role By ID
func (s *roleService) FindByID(
	ctx context.Context,
	id uint,
) (*roleModel.Role, error) {

	role, err := s.roleRepo.FindByID(ctx, id)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrRoleNotFound
		}

		return nil, err
	}

	return role, nil
}

// Find Role By Name
func (s *roleService) FindByName(
	ctx context.Context,
	name string,
) (*roleModel.Role, error) {

	name = strings.TrimSpace(name)

	if name == "" {
		return nil, apperrors.ErrInvalidRoleName
	}

	role, err := s.roleRepo.FindByName(ctx, name)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrRoleNotFound
		}

		return nil, err
	}

	return role, nil
}

// Update Role
func (s *roleService) Update(
	ctx context.Context,
	id uint,
	name string,
	description string,
) (*roleModel.Role, error) {

	name = strings.TrimSpace(name)
	description = strings.TrimSpace(description)

	if name == "" {
		return nil, apperrors.ErrInvalidRoleName
	}

	role, err := s.roleRepo.FindByID(ctx, id)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrRoleNotFound
		}

		return nil, err
	}

	// Check whether another role already uses this name.
	existing, err := s.roleRepo.FindByName(ctx, name)

	if err == nil && existing != nil && existing.ID != role.ID {
		return nil, apperrors.ErrRoleAlreadyExists
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	role.Name = name
	role.Description = description

	if err := s.roleRepo.Update(ctx, role); err != nil {
		return nil, err
	}

	return role, nil
}

// Delete Role
func (s *roleService) Delete(
	ctx context.Context,
	id uint,
) error {

	role, err := s.roleRepo.FindByID(ctx, id)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.ErrRoleNotFound
		}

		return err
	}

	return s.roleRepo.Delete(ctx, role.ID)
}
