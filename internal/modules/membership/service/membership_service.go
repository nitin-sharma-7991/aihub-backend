package service

import (
	"context"
	"errors"
	"time"

	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/membership/dto"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/membership/model"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/membership/repository"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/apperrors"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/pagination"

	"gorm.io/gorm"
)

type MembershipService interface {
	Create(
		ctx context.Context,
		organizationID uint,
		req dto.CreateMembershipRequest,
	) (*dto.MembershipResponse, error)

	FindAll(
		ctx context.Context,
		req pagination.Request,
	) ([]dto.MembershipResponse, pagination.Meta, error)

	FindByID(
		ctx context.Context,
		id uint,
	) (*dto.MembershipResponse, error)

	FindByIDAndOrganizationID(
		ctx context.Context,
		id uint,
		organizationID uint,
	) (*dto.MembershipResponse, error)

	FindAllByOrganizationID(
		ctx context.Context,
		organizationID uint,
		req pagination.Request,
	) ([]dto.MembershipResponse, pagination.Meta, error)

	Update(
		ctx context.Context,
		id uint,
		req dto.UpdateMembershipRequest,
	) (*dto.MembershipResponse, error)

	UpdateByOrganizationID(
		ctx context.Context,
		id uint,
		organizationID uint,
		req dto.UpdateMembershipRequest,
	) (*dto.MembershipResponse, error)

	Delete(
		ctx context.Context,
		id uint,
	) error

	DeleteByOrganizationID(
		ctx context.Context,
		id uint,
		organizationID uint,
	) error
}

type membershipService struct {
	membershipRepo repository.MembershipRepository
}

func NewMembershipService(
	repo repository.MembershipRepository,
) MembershipService {

	return &membershipService{
		membershipRepo: repo,
	}
}

// ================================================================
// CREATE
// ================================================================

func (s *membershipService) Create(
	ctx context.Context,
	organizationID uint,
	req dto.CreateMembershipRequest,
) (*dto.MembershipResponse, error) {

	existing, err := s.membershipRepo.FindByUserAndOrganization(
		ctx,
		req.UserID,
		organizationID,
	)

	if err == nil && existing != nil {
		return nil, apperrors.ErrMembershipAlreadyExists
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	membership := &model.Membership{
		OrganizationID: organizationID,
		UserID:         req.UserID,
		RoleID:         req.RoleID,
		JoinedAt:       time.Now(),
	}

	if err := s.membershipRepo.Create(
		ctx,
		membership,
	); err != nil {
		return nil, err
	}

	return toMembershipResponse(membership), nil
}

// ================================================================
// FIND ALL
// ================================================================

func (s *membershipService) FindAll(
	ctx context.Context,
	req pagination.Request,
) ([]dto.MembershipResponse, pagination.Meta, error) {

	req.Normalize()

	memberships, total, err := s.membershipRepo.FindAll(
		ctx,
		req,
	)
	if err != nil {
		return nil, pagination.Meta{}, err
	}

	response := make(
		[]dto.MembershipResponse,
		0,
		len(memberships),
	)

	for _, membership := range memberships {
		response = append(
			response,
			*toMembershipResponse(&membership),
		)
	}

	meta := pagination.NewMeta(
		req,
		total,
	)

	return response, meta, nil
}

// ================================================================
// FIND ALL BY ORGANIZATION
// ================================================================

func (s *membershipService) FindAllByOrganizationID(
	ctx context.Context,
	organizationID uint,
	req pagination.Request,
) ([]dto.MembershipResponse, pagination.Meta, error) {

	req.Normalize()

	memberships, total, err :=
		s.membershipRepo.FindAllByOrganizationID(
			ctx,
			organizationID,
			req,
		)

	if err != nil {
		return nil, pagination.Meta{}, err
	}

	response := make(
		[]dto.MembershipResponse,
		0,
		len(memberships),
	)

	for _, membership := range memberships {
		response = append(
			response,
			*toMembershipResponse(&membership),
		)
	}

	meta := pagination.NewMeta(
		req,
		total,
	)

	return response, meta, nil
}

// ================================================================
// FIND BY ID
// ================================================================

func (s *membershipService) FindByID(
	ctx context.Context,
	id uint,
) (*dto.MembershipResponse, error) {

	membership, err := s.membershipRepo.FindByID(
		ctx,
		id,
	)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrMembershipNotFound
		}

		return nil, err
	}

	return toMembershipResponse(membership), nil
}

// ================================================================
// FIND BY ID + ORGANIZATION
// ================================================================

func (s *membershipService) FindByIDAndOrganizationID(
	ctx context.Context,
	id uint,
	organizationID uint,
) (*dto.MembershipResponse, error) {

	membership, err :=
		s.membershipRepo.FindByIDAndOrganizationID(
			ctx,
			id,
			organizationID,
		)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrMembershipNotFound
		}

		return nil, err
	}

	return toMembershipResponse(membership), nil
}

// ================================================================
// UPDATE
// ================================================================

func (s *membershipService) Update(
	ctx context.Context,
	id uint,
	req dto.UpdateMembershipRequest,
) (*dto.MembershipResponse, error) {

	membership, err := s.membershipRepo.FindByID(
		ctx,
		id,
	)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrMembershipNotFound
		}

		return nil, err
	}

	membership.RoleID = req.RoleID

	if err := s.membershipRepo.Update(
		ctx,
		membership,
	); err != nil {
		return nil, err
	}

	return toMembershipResponse(membership), nil
}

// ================================================================
// UPDATE BY ORGANIZATION
// ================================================================

func (s *membershipService) UpdateByOrganizationID(
	ctx context.Context,
	id uint,
	organizationID uint,
	req dto.UpdateMembershipRequest,
) (*dto.MembershipResponse, error) {

	membership, err :=
		s.membershipRepo.FindByIDAndOrganizationID(
			ctx,
			id,
			organizationID,
		)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrMembershipNotFound
		}

		return nil, err
	}

	membership.RoleID = req.RoleID

	if err := s.membershipRepo.Update(
		ctx,
		membership,
	); err != nil {
		return nil, err
	}

	return toMembershipResponse(membership), nil
}

// ================================================================
// DELETE
// ================================================================

func (s *membershipService) Delete(
	ctx context.Context,
	id uint,
) error {

	membership, err := s.membershipRepo.FindByID(
		ctx,
		id,
	)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.ErrMembershipNotFound
		}

		return err
	}

	return s.membershipRepo.Delete(
		ctx,
		membership.ID,
	)
}

// ================================================================
// DELETE BY ORGANIZATION
// ================================================================

func (s *membershipService) DeleteByOrganizationID(
	ctx context.Context,
	id uint,
	organizationID uint,
) error {

	membership, err :=
		s.membershipRepo.FindByIDAndOrganizationID(
			ctx,
			id,
			organizationID,
		)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.ErrMembershipNotFound
		}

		return err
	}

	return s.membershipRepo.Delete(
		ctx,
		membership.ID,
	)
}

// ================================================================
// RESPONSE MAPPER
// ================================================================

func toMembershipResponse(
	membership *model.Membership,
) *dto.MembershipResponse {

	return &dto.MembershipResponse{
		ID:             membership.ID,
		OrganizationID: membership.OrganizationID,
		UserID:         membership.UserID,
		RoleID:         membership.RoleID,
		InvitedBy:      membership.InvitedBy,
	}
}
