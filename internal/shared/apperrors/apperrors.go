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

	// Role
	ErrRoleNotFound      = errors.New("role not found")
	ErrRoleAlreadyExists = errors.New("role already exists")
	ErrInvalidRoleName   = errors.New("invalid role")

	// Permission
	ErrPermissionNotFound      = errors.New("permission not found")
	ErrPermissionAlreadyExists = errors.New("permission already exists")
	ErrInvalidPermissionName   = errors.New("invalid permission")

	// Role Permission
	ErrRolePermissionNotFound      = errors.New("role permission assignment not found")
	ErrRolePermissionAlreadyExists = errors.New("permission is already assigned to this role")

	// Common
	ErrInternalServer = errors.New("internal server error")
)
