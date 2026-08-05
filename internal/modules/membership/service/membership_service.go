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

	Update(
		ctx context.Context,
		id uint,
		req dto.UpdateMembershipRequest,
	) (*dto.MembershipResponse, error)

	Delete(
		ctx context.Context,
		id uint,
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

func (s *membershipService) Create(
	ctx context.Context,
	req dto.CreateMembershipRequest,
) (*dto.MembershipResponse, error) {

	existing, err := s.membershipRepo.FindByUserAndOrganization(
		ctx,
		req.UserID,
		req.OrganizationID,
	)

	if err == nil && existing != nil {
		return nil, apperrors.ErrMembershipAlreadyExists
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	membership := &model.Membership{
		OrganizationID: req.OrganizationID,
		UserID:         req.UserID,
		Role:           req.Role,
		JoinedAt:       time.Now(),
	}

	if err := s.membershipRepo.Create(ctx, membership); err != nil {
		return nil, err
	}

	return toMembershipResponse(membership), nil
}

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

	response := make([]dto.MembershipResponse, 0, len(memberships))

	for _, membership := range memberships {
		response = append(response, dto.MembershipResponse{
			ID:             membership.ID,
			UserID:         membership.UserID,
			OrganizationID: membership.OrganizationID,
			Role:           membership.Role,
		})
	}

	meta := pagination.NewMeta(req, total)

	return response, meta, nil
}

func (s *membershipService) FindByID(
	ctx context.Context,
	id uint,
) (*dto.MembershipResponse, error) {

	membership, err := s.membershipRepo.FindByID(ctx, id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrMembershipNotFound
		}

		return nil, err
	}

	return toMembershipResponse(membership), nil
}

func (s *membershipService) Update(
	ctx context.Context,
	id uint,
	req dto.UpdateMembershipRequest,
) (*dto.MembershipResponse, error) {

	membership, err := s.membershipRepo.FindByID(ctx, id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrMembershipNotFound
		}

		return nil, err
	}

	membership.Role = req.Role

	if err := s.membershipRepo.Update(ctx, membership); err != nil {
		return nil, err
	}

	return toMembershipResponse(membership), nil
}

func (s *membershipService) Delete(
	ctx context.Context,
	id uint,
) error {

	membership, err := s.membershipRepo.FindByID(ctx, id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.ErrMembershipNotFound
		}

		return err
	}

	return s.membershipRepo.Delete(ctx, membership.ID)
}

func toMembershipResponse(
	membership *model.Membership,
) *dto.MembershipResponse {

	return &dto.MembershipResponse{
		ID:             membership.ID,
		OrganizationID: membership.OrganizationID,
		UserID:         membership.UserID,
		Role:           membership.Role,
		InvitedBy:      membership.InvitedBy,
	}
}
