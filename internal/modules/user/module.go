package user

import (
	"gorm.io/gorm"

	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/user/handler"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/user/repository"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/user/service"
)

type Module struct {
	Handler *handler.UserHandler
	Service service.UserService
	Repo    repository.UserRepository
}

func New(db *gorm.DB) *Module {

	repo := repository.NewUserRepository(db)

	svc := service.NewUserService(repo)

	h := handler.NewUserHandler(svc)

	return &Module{
		Handler: h,
		Service: svc,
		Repo:    repo,
	}
}
