package service

import (
	"context"
	"errors"

	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/user/dto"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/user/model"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/user/repository"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/apperrors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Create(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error)

	GetByID(ctx context.Context, id uint) (*dto.UserResponse, error)

	Update(ctx context.Context, id uint, req dto.UpdateUserRequest) (*dto.UserResponse, error)

	Delete(ctx context.Context, id uint) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepo: repo,
	}
}

func (s *userService) Create(
	ctx context.Context,
	req dto.CreateUserRequest,
) (*dto.UserResponse, error) {

	// Check duplicate email
	existingUser, err := s.userRepo.FindByEmail(ctx, req.Email)

	if err == nil && existingUser != nil {
		return nil, apperrors.ErrEmailAlreadyExists
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return toUserResponse(user), nil
}

func (s *userService) GetByID(
	ctx context.Context,
	id uint,
) (*dto.UserResponse, error) {

	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrUserNotFound
		}

		return nil, err
	}

	return toUserResponse(user), nil
}

func (s *userService) Update(
	ctx context.Context,
	id uint,
	req dto.UpdateUserRequest,
) (*dto.UserResponse, error) {

	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrUserNotFound
		}

		return nil, err
	}

	user.Name = req.Name

	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return toUserResponse(user), nil
}

func (s *userService) Delete(
	ctx context.Context,
	id uint,
) error {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.ErrUserNotFound
		}

		return err
	}

	return s.userRepo.Delete(ctx, user.ID)
}

//helper
func toUserResponse(user *model.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
