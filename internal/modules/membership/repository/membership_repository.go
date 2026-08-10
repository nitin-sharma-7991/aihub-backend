package repository

import (
	"context"

	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/membership/model"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/pagination"
	"gorm.io/gorm"
)

type MembershipRepository interface {
	Create(ctx context.Context, membership *model.Membership) error

	FindAll(
		ctx context.Context,
		req pagination.Request,
	) ([]model.Membership, int64, error)

	FindByID(ctx context.Context, id uint) (*model.Membership, error)

	FindByUserAndOrganization(
		ctx context.Context,
		userID uint,
		organizationID uint,
	) (*model.Membership, error)

	Update(ctx context.Context, membership *model.Membership) error

	Delete(ctx context.Context, id uint) error
}

type membershipRepository struct {
	db *gorm.DB
}

func NewMembershipRepository(db *gorm.DB) MembershipRepository {
	return &membershipRepository{
		db: db,
	}
}

// Create Membership
func (r *membershipRepository) Create(
	ctx context.Context,
	membership *model.Membership,
) error {

	return r.db.
		WithContext(ctx).
		Create(membership).
		Error
}

// Find All Memberships
func (r *membershipRepository) FindAll(
	ctx context.Context,
	req pagination.Request,
) ([]model.Membership, int64, error) {

	var (
		memberships []model.Membership
		total       int64
	)

	db := r.db.WithContext(ctx)

	if err := db.Model(&model.Membership{}).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := db.
		Preload("Role").
		Limit(req.Limit).
		Offset(req.Offset()).
		Find(&memberships).Error; err != nil {
		return nil, 0, err
	}

	return memberships, total, nil
}

// Find Membership By ID
func (r *membershipRepository) FindByID(
	ctx context.Context,
	id uint,
) (*model.Membership, error) {

	var membership model.Membership

	err := r.db.
		WithContext(ctx).
		Preload("Role").
		First(&membership, id).
		Error

	if err != nil {
		return nil, err
	}

	return &membership, nil
}

// Find Membership By User + Organization
func (r *membershipRepository) FindByUserAndOrganization(
	ctx context.Context,
	userID uint,
	organizationID uint,
) (*model.Membership, error) {

	var membership model.Membership

	err := r.db.
		WithContext(ctx).
		Preload("Role").
		Where(
			"user_id = ? AND organization_id = ?",
			userID,
			organizationID,
		).
		First(&membership).
		Error

	if err != nil {
		return nil, err
	}

	return &membership, nil
}

// Update Membership
func (r *membershipRepository) Update(
	ctx context.Context,
	membership *model.Membership,
) error {

	return r.db.
		WithContext(ctx).
		Save(membership).
		Error
}

// Delete Membership
func (r *membershipRepository) Delete(
	ctx context.Context,
	id uint,
) error {

	return r.db.
		WithContext(ctx).
		Delete(&model.Membership{}, id).
		Error
}
