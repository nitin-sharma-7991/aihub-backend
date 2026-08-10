package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/role_permission/dto"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/role_permission/service"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/apperrors"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/response"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/validation"
)

type RolePermissionHandler struct {
	rolePermissionService service.RolePermissionService
}

func NewRolePermissionHandler(
	rolePermissionService service.RolePermissionService,
) *RolePermissionHandler {
	return &RolePermissionHandler{
		rolePermissionService: rolePermissionService,
	}
}

// POST /roles/:role_id/permissions
func (h *RolePermissionHandler) AssignPermission(
	ctx *gin.Context,
) {
	roleID, err := strconv.ParseUint(
		ctx.Param("role_id"),
		10,
		64,
	)

	if err != nil {
		response.BadRequest(
			ctx,
			"Invalid role id",
			nil,
		)
		return
	}

	var req dto.AssignPermissionRequest

	if err := validation.BindJSON(ctx, &req); err != nil {
		return
	}

	rolePermission, err := h.rolePermissionService.AssignPermission(
		ctx.Request.Context(),
		uint(roleID),
		req.PermissionID,
	)

	if err != nil {

		switch {
		case errors.Is(err, apperrors.ErrRoleNotFound):
			response.NotFound(ctx, err.Error())

		case errors.Is(err, apperrors.ErrPermissionNotFound):
			response.NotFound(ctx, err.Error())

		case errors.Is(err, apperrors.ErrRolePermissionAlreadyExists):
			response.Conflict(ctx, err.Error())

		default:
			response.InternalServerError(ctx)
		}

		return
	}

	response.Created(
		ctx,
		"Permission assigned to role successfully",
		rolePermission,
	)
}

// GET /roles/:role_id/permissions
func (h *RolePermissionHandler) GetPermissions(
	ctx *gin.Context,
) {
	roleID, err := strconv.ParseUint(
		ctx.Param("role_id"),
		10,
		64,
	)

	if err != nil {
		response.BadRequest(
			ctx,
			"Invalid role id",
			nil,
		)
		return
	}

	permissions, err := h.rolePermissionService.FindPermissionsByRoleID(
		ctx.Request.Context(),
		uint(roleID),
	)

	if err != nil {

		if errors.Is(err, apperrors.ErrRoleNotFound) {
			response.NotFound(ctx, err.Error())
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Success(
		ctx,
		"Role permissions fetched successfully",
		permissions,
	)
}

// DELETE /roles/:role_id/permissions/:permission_id
func (h *RolePermissionHandler) RevokePermission(
	ctx *gin.Context,
) {
	roleID, err := strconv.ParseUint(
		ctx.Param("role_id"),
		10,
		64,
	)

	if err != nil {
		response.BadRequest(
			ctx,
			"Invalid role id",
			nil,
		)
		return
	}

	permissionID, err := strconv.ParseUint(
		ctx.Param("permission_id"),
		10,
		64,
	)

	if err != nil {
		response.BadRequest(
			ctx,
			"Invalid permission id",
			nil,
		)
		return
	}

	err = h.rolePermissionService.RevokePermission(
		ctx.Request.Context(),
		uint(roleID),
		uint(permissionID),
	)

	if err != nil {

		switch {
		case errors.Is(err, apperrors.ErrRoleNotFound):
			response.NotFound(ctx, err.Error())

		case errors.Is(err, apperrors.ErrRolePermissionNotFound):
			response.NotFound(ctx, err.Error())

		default:
			response.InternalServerError(ctx)
		}

		return
	}

	response.Success(
		ctx,
		"Permission revoked from role successfully",
		nil,
	)
}
