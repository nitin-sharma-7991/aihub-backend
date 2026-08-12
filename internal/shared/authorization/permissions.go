package authorization

const (
	// User permissions
	UserCreate = "user.create"
	UserRead   = "user.read"
	UserUpdate = "user.update"
	UserDelete = "user.delete"

	// Organization permissions
	OrganizationCreate = "organization.create"
	OrganizationRead   = "organization.read"
	OrganizationUpdate = "organization.update"
	OrganizationDelete = "organization.delete"

	// Membership permissions
	MembershipCreate = "membership.create"
	MembershipRead   = "membership.read"
	MembershipUpdate = "membership.update"
	MembershipDelete = "membership.delete"

	// Role permissions
	RoleCreate = "role.create"
	RoleRead   = "role.read"
	RoleUpdate = "role.update"
	RoleDelete = "role.delete"

	// Permission management
	PermissionCreate = "permission.create"
	PermissionRead   = "permission.read"
	PermissionUpdate = "permission.update"
	PermissionDelete = "permission.delete"

	// Role permission management
	RolePermissionAssign = "role_permission.assign"
	RolePermissionRevoke = "role_permission.revoke"
)
