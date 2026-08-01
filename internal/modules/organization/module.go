package organization

import (
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/organization/handler"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/organization/repository"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/organization/service"
	"gorm.io/gorm"
)

type Module struct {
	Handler *handler.OrganizationHandler
	Service service.OrganizationService
	OrgRepo repository.OrganizationRepository
}

func New(db *gorm.DB) *Module {

	repo := repository.NewOrganizationRepository(db)

	svc := service.NewOrganizationService(repo)

	h := handler.NewOrganizationHandler(svc)

	return &Module{
		Handler: h,
		Service: svc,
		OrgRepo: repo,
	}
}
