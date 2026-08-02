package repository

import (
	"context"

	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/membership/model"
	"gorm.io/gorm"
)

type MembershipRepository interface {
	Create(ctx context.Context, membership *model.Membership) error

	FindAll(ctx context.Context) ([]model.Membership, error)

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
) ([]model.Membership, error) {

	var memberships []model.Membership

	err := r.db.
		WithContext(ctx).
		Find(&memberships).
		Error

	if err != nil {
		return nil, err
	}

	return memberships, nil
}

// Find Membership By ID
func (r *membershipRepository) FindByID(
	ctx context.Context,
	id uint,
) (*model.Membership, error) {

	var membership model.Membership

	err := r.db.
		WithContext(ctx).
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
