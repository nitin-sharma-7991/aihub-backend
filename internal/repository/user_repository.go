package repository

import (
	"context"

	"github.com/nitin-sharma-7991/aihub-backend/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error

	FindByID(ctx context.Context, id uint) (*model.User, error)

	FindByEmail(ctx context.Context, email string) (*model.User, error)

	Update(ctx context.Context, user *model.User) error

	Delete(ctx context.Context, id uint) error
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// Create inserts a new user into the database.
func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).
		Create(user).Error
}

// FindByID retrieves a user by its ID.
func (r *userRepository) FindByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User

	err := r.db.WithContext(ctx).
		First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByEmail retrieves a user by email.
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update updates an existing user.
func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).
		Save(user).Error
}

// Delete removes a user by ID.
func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Delete(&model.User{}, id).Error
}
