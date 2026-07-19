package service

import (
	"context"
	"errors"

	"github.com/nitin-sharma-7991/aihub-backend/internal/apperrors"
	"github.com/nitin-sharma-7991/aihub-backend/internal/dto"
	"github.com/nitin-sharma-7991/aihub-backend/internal/model"
	"github.com/nitin-sharma-7991/aihub-backend/internal/repository"
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
	repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repository: repo,
	}
}

func (s *userService) Create(
	ctx context.Context,
	req dto.CreateUserRequest,
) (*dto.UserResponse, error) {

	// Check duplicate email
	existingUser, err := s.repository.FindByEmail(ctx, req.Email)

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

	if err := s.repository.Create(ctx, user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *userService) GetByID(
	ctx context.Context,
	id uint,
) (*dto.UserResponse, error) {

	user, err := s.repository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrUserNotFound
		}

		return nil, err
	}

	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *userService) Update(
	ctx context.Context,
	id uint,
	req dto.UpdateUserRequest,
) (*dto.UserResponse, error) {

	user, err := s.repository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrUserNotFound
		}

		return nil, err
	}

	user.Name = req.Name

	if err := s.repository.Update(ctx, user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *userService) Delete(
	ctx context.Context,
	id uint,
) error {
	user, err := s.repository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.ErrUserNotFound
		}

		return err
	}

	return s.repository.Delete(ctx, user.ID)
}
