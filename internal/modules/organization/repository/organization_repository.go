package repository

import (
	"context"

	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/organization/model"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	Create(ctx context.Context, org *model.Organization) error

	FindAll(ctx context.Context) ([]model.Organization, error)

	FindByID(ctx context.Context, id uint) (*model.Organization, error)

	FindBySlug(ctx context.Context, slug string) (*model.Organization, error)

	Update(ctx context.Context, org *model.Organization) error

	Delete(ctx context.Context, id uint) error
}

type organizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepository{
		db: db,
	}
}

// Create Organization
func (r *organizationRepository) Create(
	ctx context.Context,
	org *model.Organization,
) error {

	return r.db.
		WithContext(ctx).
		Create(org).
		Error
}

// Find All Organization
func (r *organizationRepository) FindAll(
	ctx context.Context,
) ([]model.Organization, error) {

	var organizations []model.Organization

	err := r.db.
		WithContext(ctx).
		Find(&organizations).
		Error

	if err != nil {
		return nil, err
	}

	return organizations, nil
}

// Find Organization By ID
func (r *organizationRepository) FindByID(
	ctx context.Context,
	id uint,
) (*model.Organization, error) {

	var organization model.Organization

	err := r.db.
		WithContext(ctx).
		First(&organization, id).
		Error

	if err != nil {
		return nil, err
	}

	return &organization, nil
}

// Find Organization By Slug
func (r *organizationRepository) FindBySlug(
	ctx context.Context,
	slug string,
) (*model.Organization, error) {

	var organization model.Organization

	err := r.db.
		WithContext(ctx).
		Where("slug = ?", slug).
		First(&organization).
		Error

	if err != nil {
		return nil, err
	}

	return &organization, nil
}

// Update Organization
func (r *organizationRepository) Update(
	ctx context.Context,
	org *model.Organization,
) error {

	return r.db.
		WithContext(ctx).
		Save(org).
		Error
}

// Delete Organization
func (r *organizationRepository) Delete(
	ctx context.Context,
	id uint,
) error {

	return r.db.
		WithContext(ctx).
		Delete(&model.Organization{}, id).
		Error
}
