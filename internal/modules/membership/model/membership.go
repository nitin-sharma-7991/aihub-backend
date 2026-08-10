package model

import (
	"time"

	organizationModel "github.com/nitin-sharma-7991/aihub-backend/internal/modules/organization/model"
	roleModel "github.com/nitin-sharma-7991/aihub-backend/internal/modules/role/model"
	userModel "github.com/nitin-sharma-7991/aihub-backend/internal/modules/user/model"
)

type Membership struct {
	ID uint `gorm:"primaryKey"`

	// One user can belong to multiple organizations.
	// One organization can have multiple users.
	// But a user can exist only once per organization.
	OrganizationID uint                            `gorm:"not null;uniqueIndex:idx_org_user"`
	Organization   *organizationModel.Organization `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	UserID uint            `gorm:"not null;uniqueIndex:idx_org_user"`
	User   *userModel.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	RoleID uint            `gorm:"not null;index"`
	Role   *roleModel.Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	InvitedBy uint
	Inviter   *userModel.User `gorm:"foreignKey:InvitedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	JoinedAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}
