package dto

type AssignPermissionRequest struct {
	PermissionID uint `json:"permission_id" binding:"required"`
}

type RolePermissionResponse struct {
	RoleID       uint `json:"role_id"`
	PermissionID uint `json:"permission_id"`
}

type PermissionResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
