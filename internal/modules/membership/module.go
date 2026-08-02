package membership

import (
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/membership/handler"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/membership/repository"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/membership/service"
	"gorm.io/gorm"
)

type Module struct {
	Handler        *handler.MembershipHandler
	Service        service.MembershipService
	MembershipRepo repository.MembershipRepository
}

func New(db *gorm.DB) *Module {
	repo := repository.NewMembershipRepository(db)

	svc := service.NewMembershipService(repo)

	h := handler.NewMembershipHandler(svc)

	return &Module{
		Handler:        h,
		Service:        svc,
		MembershipRepo: repo,
	}
}
