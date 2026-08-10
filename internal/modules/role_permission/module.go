package role_permission

import (
	permissionRepository "github.com/nitin-sharma-7991/aihub-backend/internal/modules/permission/repository"
	roleRepository "github.com/nitin-sharma-7991/aihub-backend/internal/modules/role/repository"
	rolePermissionRepository "github.com/nitin-sharma-7991/aihub-backend/internal/modules/role_permission/repository"

	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/role_permission/handler"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/role_permission/service"

	"gorm.io/gorm"
)

type Module struct {
	Handler            *handler.RolePermissionHandler
	Service            service.RolePermissionService
	RolePermissionRepo rolePermissionRepository.RolePermissionRepository
}

func New(db *gorm.DB) *Module {

	rolePermissionRepo := rolePermissionRepository.NewRolePermissionRepository(db)

	permissionRepo := permissionRepository.NewPermissionRepository(db)

	roleRepo := roleRepository.NewRoleRepository(db)

	svc := service.NewRolePermissionService(
		rolePermissionRepo,
		permissionRepo,
		roleRepo,
	)

	h := handler.NewRolePermissionHandler(svc)

	return &Module{
		Handler:            h,
		Service:            svc,
		RolePermissionRepo: rolePermissionRepo,
	}
}
