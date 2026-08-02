package apperrors

import "errors"

var (
	// User
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")

	// Auth
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrInvalidToken       = errors.New("invalid token")
	ErrExpiredToken       = errors.New("token has expired")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")
	ErrMissingToken       = errors.New("authorization token is required")

	// Organization
	ErrOrganizationAlreadyExists = errors.New("organization already exists")
	ErrOrganizationNotFound      = errors.New("organization not found")

	// Membership
	ErrMembershipNotFound      = errors.New("membership not found")
	ErrMembershipAlreadyExists = errors.New("membership already exists")

	// Common
	ErrInternalServer = errors.New("internal server error")
)
