package permission

import (
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/permission/handler"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/permission/repository"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/permission/service"

	"gorm.io/gorm"
)

type Module struct {
	Handler        *handler.PermissionHandler
	Service        service.PermissionService
	PermissionRepo repository.PermissionRepository
}

func New(db *gorm.DB) *Module {

	repo := repository.NewPermissionRepository(db)

	svc := service.NewPermissionService(repo)

	h := handler.NewPermissionHandler(svc)

	return &Module{
		Handler:        h,
		Service:        svc,
		PermissionRepo: repo,
	}
}
