package auth

import (
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/auth/handler"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/auth/service"

	userRepository "github.com/nitin-sharma-7991/aihub-backend/internal/modules/user/repository"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/config"
)

type Module struct {
	Handler *handler.AuthHandler
	Service service.AuthService
}

func New(
	userRepo userRepository.UserRepository,
	cfg *config.Config,
) *Module {

	svc := service.NewAuthService(
		userRepo,
		cfg,
	)

	h := handler.NewAuthHandler(svc)

	return &Module{
		Handler: h,
		Service: svc,
	}
}
