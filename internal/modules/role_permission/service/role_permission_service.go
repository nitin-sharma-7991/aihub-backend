package service

import (
	"context"
	"errors"

	permissionModel "github.com/nitin-sharma-7991/aihub-backend/internal/modules/permission/model"
	permissionRepository "github.com/nitin-sharma-7991/aihub-backend/internal/modules/permission/repository"
	roleRepository "github.com/nitin-sharma-7991/aihub-backend/internal/modules/role/repository"
	rolePermissionModel "github.com/nitin-sharma-7991/aihub-backend/internal/modules/role_permission/model"
	rolePermissionRepository "github.com/nitin-sharma-7991/aihub-backend/internal/modules/role_permission/repository"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/apperrors"

	"gorm.io/gorm"
)

type RolePermissionService interface {
	AssignPermission(
		ctx context.Context,
		roleID uint,
		permissionID uint,
	) (*rolePermissionModel.RolePermission, error)

	FindPermissionsByRoleID(
		ctx context.Context,
		roleID uint,
	) ([]permissionModel.Permission, error)

	RevokePermission(
		ctx context.Context,
		roleID uint,
		permissionID uint,
	) error
}

type rolePermissionService struct {
	rolePermissionRepo rolePermissionRepository.RolePermissionRepository
	permissionRepo     permissionRepository.PermissionRepository
	roleRepo           roleRepository.RoleRepository
}

func NewRolePermissionService(
	rolePermissionRepo rolePermissionRepository.RolePermissionRepository,
	permissionRepo permissionRepository.PermissionRepository,
	roleRepo roleRepository.RoleRepository,
) RolePermissionService {
	return &rolePermissionService{
		rolePermissionRepo: rolePermissionRepo,
		permissionRepo:     permissionRepo,
		roleRepo:           roleRepo,
	}
}

// Assign Permission To Role
func (s *rolePermissionService) AssignPermission(
	ctx context.Context,
	roleID uint,
	permissionID uint,
) (*rolePermissionModel.RolePermission, error) {

	// Make sure role exists.
	_, err := s.roleRepo.FindByID(ctx, roleID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrRoleNotFound
		}

		return nil, err
	}

	// Make sure permission exists.
	_, err = s.permissionRepo.FindByID(ctx, permissionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrPermissionNotFound
		}

		return nil, err
	}

	// Prevent duplicate assignment.
	exists, err := s.rolePermissionRepo.Exists(
		ctx,
		roleID,
		permissionID,
	)

	if err != nil {
		return nil, err
	}

	if exists {
		return nil, apperrors.ErrRolePermissionAlreadyExists
	}

	rolePermission := &rolePermissionModel.RolePermission{
		RoleID:       roleID,
		PermissionID: permissionID,
	}

	if err := s.rolePermissionRepo.Create(
		ctx,
		rolePermission,
	); err != nil {
		return nil, err
	}

	return rolePermission, nil
}

// Find Permissions Assigned To Role
func (s *rolePermissionService) FindPermissionsByRoleID(
	ctx context.Context,
	roleID uint,
) ([]permissionModel.Permission, error) {

	// Make sure role exists.
	_, err := s.roleRepo.FindByID(ctx, roleID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrRoleNotFound
		}

		return nil, err
	}

	return s.permissionRepo.FindByRoleID(
		ctx,
		roleID,
	)
}

// Revoke Permission From Role
func (s *rolePermissionService) RevokePermission(
	ctx context.Context,
	roleID uint,
	permissionID uint,
) error {

	// Make sure role exists.
	_, err := s.roleRepo.FindByID(ctx, roleID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.ErrRoleNotFound
		}

		return err
	}

	// Make sure assignment exists.
	exists, err := s.rolePermissionRepo.Exists(
		ctx,
		roleID,
		permissionID,
	)

	if err != nil {
		return err
	}

	if !exists {
		return apperrors.ErrRolePermissionNotFound
	}

	return s.rolePermissionRepo.Delete(
		ctx,
		roleID,
		permissionID,
	)
}
