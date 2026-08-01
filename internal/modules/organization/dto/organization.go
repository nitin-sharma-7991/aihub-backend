package dto

type CreateOrganizationRequest struct {
	Name        string `json:"name" binding:"required,min=3,max=100"`
	Slug        string `json:"slug" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"required,min=8"`
}

type OrganizationResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

type UpdateOrganizationRequest struct {
	Name        string `json:"name" binding:"required,min=3,max=100"`
	Slug        string `json:"slug" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"required,min=8"`
}
