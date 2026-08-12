package dto

type CreateMembershipRequest struct {
	UserID uint `json:"user_id" binding:"required"`
	RoleID uint `json:"role_id" binding:"required"`
}

type MembershipResponse struct {
	ID             uint `json:"id"`
	OrganizationID uint `json:"organization_id"`
	UserID         uint `json:"user_id"`
	RoleID         uint `json:"role_id"`
	InvitedBy      uint `json:"invited_by"`
}

type UpdateMembershipRequest struct {
	RoleID uint `json:"role_id" binding:"required"`
}
