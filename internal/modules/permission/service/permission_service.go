package service

import (
	"context"
	"errors"
	"strings"

	permissionModel "github.com/nitin-sharma-7991/aihub-backend/internal/modules/permission/model"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/permission/repository"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/apperrors"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/pagination"

	"gorm.io/gorm"
)

type PermissionService interface {
	Create(
		ctx context.Context,
		name string,
		description string,
	) (*permissionModel.Permission, error)

	FindAll(
		ctx context.Context,
		req pagination.Request,
	) ([]permissionModel.Permission, pagination.Meta, error)

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
		id uint,
		name string,
		description string,
	) (*permissionModel.Permission, error)

	Delete(
		ctx context.Context,
		id uint,
	) error

	FindByRoleID(
		ctx context.Context,
		roleID uint,
	) ([]permissionModel.Permission, error)
}

type permissionService struct {
	permissionRepo repository.PermissionRepository
}

func NewPermissionService(
	permissionRepo repository.PermissionRepository,
) PermissionService {
	return &permissionService{
		permissionRepo: permissionRepo,
	}
}

// Create Permission
func (s *permissionService) Create(
	ctx context.Context,
	name string,
	description string,
) (*permissionModel.Permission, error) {

	name = strings.TrimSpace(name)
	description = strings.TrimSpace(description)

	if name == "" {
		return nil, apperrors.ErrInvalidPermissionName
	}

	// Prevent duplicate permission names.
	existing, err := s.permissionRepo.FindByName(ctx, name)

	if err == nil && existing != nil {
		return nil, apperrors.ErrPermissionAlreadyExists
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	permission := &permissionModel.Permission{
		Name:        name,
		Description: description,
	}

	if err := s.permissionRepo.Create(ctx, permission); err != nil {
		return nil, err
	}

	return permission, nil
}

// Find All Permissions
func (s *permissionService) FindAll(
	ctx context.Context,
	req pagination.Request,
) ([]permissionModel.Permission, pagination.Meta, error) {

	req.Normalize()

	permissions, total, err := s.permissionRepo.FindAll(
		ctx,
		req,
	)

	if err != nil {
		return nil, pagination.Meta{}, err
	}

	meta := pagination.NewMeta(req, total)

	return permissions, meta, nil
}

// Find Permission By ID
func (s *permissionService) FindByID(
	ctx context.Context,
	id uint,
) (*permissionModel.Permission, error) {

	permission, err := s.permissionRepo.FindByID(ctx, id)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrPermissionNotFound
		}

		return nil, err
	}

	return permission, nil
}

// Find Permission By Name
func (s *permissionService) FindByName(
	ctx context.Context,
	name string,
) (*permissionModel.Permission, error) {

	name = strings.TrimSpace(name)

	if name == "" {
		return nil, apperrors.ErrInvalidPermissionName
	}

	permission, err := s.permissionRepo.FindByName(ctx, name)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrPermissionNotFound
		}

		return nil, err
	}

	return permission, nil
}

// Update Permission
func (s *permissionService) Update(
	ctx context.Context,
	id uint,
	name string,
	description string,
) (*permissionModel.Permission, error) {

	name = strings.TrimSpace(name)
	description = strings.TrimSpace(description)

	if name == "" {
		return nil, apperrors.ErrInvalidPermissionName
	}

	permission, err := s.permissionRepo.FindByID(ctx, id)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrPermissionNotFound
		}

		return nil, err
	}

	// Check whether another permission already uses this name.
	existing, err := s.permissionRepo.FindByName(ctx, name)

	if err == nil && existing != nil && existing.ID != permission.ID {
		return nil, apperrors.ErrPermissionAlreadyExists
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	permission.Name = name
	permission.Description = description

	if err := s.permissionRepo.Update(ctx, permission); err != nil {
		return nil, err
	}

	return permission, nil
}

// Delete Permission
func (s *permissionService) Delete(
	ctx context.Context,
	id uint,
) error {

	permission, err := s.permissionRepo.FindByID(ctx, id)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.ErrPermissionNotFound
		}

		return err
	}

	return s.permissionRepo.Delete(ctx, permission.ID)
}

// Find Permissions Assigned To Role
func (s *permissionService) FindByRoleID(
	ctx context.Context,
	roleID uint,
) ([]permissionModel.Permission, error) {

	return s.permissionRepo.FindByRoleID(
		ctx,
		roleID,
	)
}
