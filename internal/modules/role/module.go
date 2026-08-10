package role

import (
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/role/handler"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/role/repository"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/role/service"

	"gorm.io/gorm"
)

type Module struct {
	Handler  *handler.RoleHandler
	Service  service.RoleService
	RoleRepo repository.RoleRepository
}

func New(db *gorm.DB) *Module {

	repo := repository.NewRoleRepository(db)

	svc := service.NewRoleService(repo)

	h := handler.NewRoleHandler(svc)

	return &Module{
		Handler:  h,
		Service:  svc,
		RoleRepo: repo,
	}
}
