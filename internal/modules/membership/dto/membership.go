package dto

type CreateMembershipRequest struct {
	OrganizationID uint   `json:"organization_id" binding:"required"`
	UserID         uint   `json:"user_id" binding:"required"`
	Role           string `json:"role" binding:"required,oneof=owner admin developer viewer"`
}

type MembershipResponse struct {
	ID             uint   `json:"id"`
	OrganizationID uint   `json:"organization_id"`
	UserID         uint   `json:"user_id"`
	Role           string `json:"role"`
	InvitedBy      uint   `json:"invited_by"`
}

type UpdateMembershipRequest struct {
	Role string `json:"role" binding:"required,oneof=owner admin developer viewer"`
}
