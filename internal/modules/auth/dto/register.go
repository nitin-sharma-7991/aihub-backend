package dto

type RegisterRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`

	Email string `json:"email" validate:"required,email"`

	Password string `json:"password" validate:"required,min=8,max=72"`
}

type RegisterResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
