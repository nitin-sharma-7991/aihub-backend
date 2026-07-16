package service

import (
	"github.com/nitin-sharma-7991/aihub-backend/internal/dto"
)

type UserService interface {
	Create(req dto.CreateUserRequest) error

	GetByID(id uint) error

	Update(id uint) error

	Delete(id uint) error
}
