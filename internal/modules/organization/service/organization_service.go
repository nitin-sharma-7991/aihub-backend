package service

import (
	"context"
	"errors"

	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/organization/dto"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/organization/model"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/organization/repository"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/apperrors"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/pagination"
	"gorm.io/gorm"
)

type OrganizationService interface {
	Create(ctx context.Context, req dto.CreateOrganizationRequest) (*dto.OrganizationResponse, error)
	FindAll(
		ctx context.Context,
		req pagination.Request,
	) ([]dto.OrganizationResponse, pagination.Meta, error)
	FindByID(ctx context.Context, id uint) (*dto.OrganizationResponse, error)
	Update(ctx context.Context, id uint, req dto.UpdateOrganizationRequest) (*dto.OrganizationResponse, error)
	Delete(ctx context.Context, id uint) error
}

type organizationService struct {
	orgRepo repository.OrganizationRepository
}

func NewOrganizationService(repo repository.OrganizationRepository) OrganizationService {
	return &organizationService{
		orgRepo: repo,
	}
}

func (s *organizationService) Create(
	ctx context.Context,
	req dto.CreateOrganizationRequest,
) (*dto.OrganizationResponse, error) {

	/// Check duplicate organization slug
	existingOrg, err := s.orgRepo.FindBySlug(ctx, req.Slug)

	if err == nil && existingOrg != nil {
		return nil, apperrors.ErrOrganizationAlreadyExists
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	org := &model.Organization{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
	}

	if err := s.orgRepo.Create(ctx, org); err != nil {
		return nil, err
	}

	return toOrganizationResponse(org), nil
}

func (s *organizationService) FindAll(
	ctx context.Context,
	req pagination.Request,
) ([]dto.OrganizationResponse, pagination.Meta, error) {

	req.Normalize()

	orgs, total, err := s.orgRepo.FindAll(
		ctx,
		req,
	)
	if err != nil {
		return nil, pagination.Meta{}, err
	}

	response := make([]dto.OrganizationResponse, 0, len(orgs))

	for _, org := range orgs {
		response = append(response, dto.OrganizationResponse{
			ID:          org.ID,
			Name:        org.Name,
			Slug:        org.Slug,
			Description: org.Description,
		})
	}

	meta := pagination.NewMeta(req, total)

	return response, meta, nil
}

func (s *organizationService) FindByID(
	ctx context.Context,
	id uint,
) (*dto.OrganizationResponse, error) {

	org, err := s.orgRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrOrganizationNotFound
		}

		return nil, err
	}

	return toOrganizationResponse(org), nil
}

func (s *organizationService) Update(
	ctx context.Context,
	id uint,
	req dto.UpdateOrganizationRequest,
) (*dto.OrganizationResponse, error) {

	org, err := s.orgRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrOrganizationNotFound
		}

		return nil, err
	}

	org.Name = req.Name
	org.Slug = req.Slug
	org.Description = req.Description

	if err := s.orgRepo.Update(ctx, org); err != nil {
		return nil, err
	}

	return toOrganizationResponse(org), nil
}

func (s *organizationService) Delete(
	ctx context.Context,
	id uint,
) error {
	org, err := s.orgRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.ErrOrganizationNotFound
		}

		return err
	}

	return s.orgRepo.Delete(ctx, org.ID)
}

//helper
func toOrganizationResponse(org *model.Organization) *dto.OrganizationResponse {
	return &dto.OrganizationResponse{
		ID:          org.ID,
		Name:        org.Name,
		Slug:        org.Slug,
		Description: org.Description,
	}
}
