package authorization

import (
	"context"
	"errors"

	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/membership/repository"
	permissionRepository "github.com/nitin-sharma-7991/aihub-backend/internal/modules/permission/repository"

	"gorm.io/gorm"
)

type Service interface {
	Can(
		ctx context.Context,
		userID uint,
		organizationID uint,
		permissionName string,
	) (bool, error)
}

type service struct {
	membershipRepo repository.MembershipRepository
	permissionRepo permissionRepository.PermissionRepository
}

func NewService(
	membershipRepo repository.MembershipRepository,
	permissionRepo permissionRepository.PermissionRepository,
) Service {
	return &service{
		membershipRepo: membershipRepo,
		permissionRepo: permissionRepo,
	}
}

func (s *service) Can(
	ctx context.Context,
	userID uint,
	organizationID uint,
	permissionName string,
) (bool, error) {

	// Find user's membership in the organization.
	membership, err := s.membershipRepo.FindByUserAndOrganization(
		ctx,
		userID,
		organizationID,
	)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	// Get all permissions assigned to the membership's role.
	permissions, err := s.permissionRepo.FindByRoleID(
		ctx,
		membership.RoleID,
	)

	if err != nil {
		return false, err
	}

	// Check whether the required permission exists.
	for _, permission := range permissions {
		if permission.Name == permissionName {
			return true, nil
		}
	}

	return false, nil
}
