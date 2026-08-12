package service

import (
	"context"
	"errors"

	authDTO "github.com/nitin-sharma-7991/aihub-backend/internal/modules/auth/dto"
	userModel "github.com/nitin-sharma-7991/aihub-backend/internal/modules/user/model"
	userRepository "github.com/nitin-sharma-7991/aihub-backend/internal/modules/user/repository"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/apperrors"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/config"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/security"

	"gorm.io/gorm"
)

type AuthService interface {
	Login(
		ctx context.Context,
		req authDTO.LoginRequest,
	) (*authDTO.LoginResponse, error)

	Register(
		ctx context.Context,
		req authDTO.RegisterRequest,
	) (*authDTO.RegisterResponse, error)

	Me(
		ctx context.Context,
		userID uint,
	) (*authDTO.MeResponse, error)
}

type authService struct {
	userRepo userRepository.UserRepository
	cfg      *config.Config
}

func NewAuthService(
	userRepo userRepository.UserRepository,
	cfg *config.Config,
) AuthService {
	return &authService{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

// Login
func (s *authService) Login(
	ctx context.Context,
	req authDTO.LoginRequest,
) (*authDTO.LoginResponse, error) {

	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrInvalidCredentials
		}

		return nil, err
	}

	if err := security.CheckPassword(
		user.Password,
		req.Password,
	); err != nil {
		return nil, apperrors.ErrInvalidCredentials
	}

	token, err := security.GenerateJWT(
		user.ID,
		user.Role,
		s.cfg.JWT.Secret,
		s.cfg.JWT.ExpiresIn,
	)
	if err != nil {
		return nil, err
	}

	return &authDTO.LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   s.cfg.JWT.ExpiresIn.String(),
	}, nil
}

// Register
func (s *authService) Register(
	ctx context.Context,
	req authDTO.RegisterRequest,
) (*authDTO.RegisterResponse, error) {

	existingUser, err := s.userRepo.FindByEmail(
		ctx,
		req.Email,
	)

	if err == nil && existingUser != nil {
		return nil, apperrors.ErrEmailAlreadyExists
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &userModel.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return &authDTO.RegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// Me
func (s *authService) Me(
	ctx context.Context,
	userID uint,
) (*authDTO.MeResponse, error) {

	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrUserNotFound
		}

		return nil, err
	}

	return &authDTO.MeResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}
