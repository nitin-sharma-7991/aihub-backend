package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/permission/dto"
	permissionModel "github.com/nitin-sharma-7991/aihub-backend/internal/modules/permission/model"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/permission/service"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/apperrors"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/pagination"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/response"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/validation"
)

type PermissionHandler struct {
	permissionService service.PermissionService
}

func NewPermissionHandler(
	permissionService service.PermissionService,
) *PermissionHandler {
	return &PermissionHandler{
		permissionService: permissionService,
	}
}

// POST /permissions
func (h *PermissionHandler) Create(ctx *gin.Context) {

	var req dto.CreatePermissionRequest

	if err := validation.BindJSON(ctx, &req); err != nil {
		return
	}

	permission, err := h.permissionService.Create(
		ctx.Request.Context(),
		req.Name,
		req.Description,
	)

	if err != nil {

		if errors.Is(err, apperrors.ErrInvalidPermissionName) {
			response.BadRequest(
				ctx,
				err.Error(),
				nil,
			)
			return
		}

		if errors.Is(err, apperrors.ErrPermissionAlreadyExists) {
			response.Conflict(
				ctx,
				err.Error(),
			)
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Created(
		ctx,
		"Permission created successfully",
		toPermissionResponse(permission),
	)
}

// GET /permissions
func (h *PermissionHandler) GetAll(ctx *gin.Context) {

	var req pagination.Request

	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.BadRequest(
			ctx,
			"Invalid pagination parameters",
			nil,
		)
		return
	}

	permissions, meta, err := h.permissionService.FindAll(
		ctx.Request.Context(),
		req,
	)

	if err != nil {
		response.InternalServerError(ctx)
		return
	}

	responses := make(
		[]dto.PermissionResponse,
		0,
		len(permissions),
	)

	for _, permission := range permissions {
		responses = append(
			responses,
			toPermissionResponse(&permission),
		)
	}

	response.SuccessWithMeta(
		ctx,
		"Permissions fetched successfully",
		responses,
		meta,
	)
}

// GET /permissions/:id
func (h *PermissionHandler) GetByID(ctx *gin.Context) {

	id, err := strconv.ParseUint(
		ctx.Param("id"),
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

	permission, err := h.permissionService.FindByID(
		ctx.Request.Context(),
		uint(id),
	)

	if err != nil {

		if errors.Is(err, apperrors.ErrPermissionNotFound) {
			response.NotFound(
				ctx,
				err.Error(),
			)
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Success(
		ctx,
		"Permission fetched successfully",
		toPermissionResponse(permission),
	)
}

// PUT /permissions/:id
func (h *PermissionHandler) Update(ctx *gin.Context) {

	id, err := strconv.ParseUint(
		ctx.Param("id"),
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

	var req dto.UpdatePermissionRequest

	if err := validation.BindJSON(ctx, &req); err != nil {
		return
	}

	permission, err := h.permissionService.Update(
		ctx.Request.Context(),
		uint(id),
		req.Name,
		req.Description,
	)

	if err != nil {

		if errors.Is(err, apperrors.ErrPermissionNotFound) {
			response.NotFound(
				ctx,
				err.Error(),
			)
			return
		}

		if errors.Is(err, apperrors.ErrInvalidPermissionName) {
			response.BadRequest(
				ctx,
				err.Error(),
				nil,
			)
			return
		}

		if errors.Is(err, apperrors.ErrPermissionAlreadyExists) {
			response.Conflict(
				ctx,
				err.Error(),
			)
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Success(
		ctx,
		"Permission updated successfully",
		toPermissionResponse(permission),
	)
}

// DELETE /permissions/:id
func (h *PermissionHandler) Delete(ctx *gin.Context) {

	id, err := strconv.ParseUint(
		ctx.Param("id"),
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

	err = h.permissionService.Delete(
		ctx.Request.Context(),
		uint(id),
	)

	if err != nil {

		if errors.Is(err, apperrors.ErrPermissionNotFound) {
			response.NotFound(
				ctx,
				err.Error(),
			)
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Success(
		ctx,
		"Permission deleted successfully",
		nil,
	)
}

// GET /permissions/role/:role_id
func (h *PermissionHandler) GetByRoleID(ctx *gin.Context) {

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

	permissions, err := h.permissionService.FindByRoleID(
		ctx.Request.Context(),
		uint(roleID),
	)

	if err != nil {
		response.InternalServerError(ctx)
		return
	}

	responses := make(
		[]dto.PermissionResponse,
		0,
		len(permissions),
	)

	for _, permission := range permissions {
		responses = append(
			responses,
			toPermissionResponse(&permission),
		)
	}

	response.Success(
		ctx,
		"Role permissions fetched successfully",
		responses,
	)
}

func toPermissionResponse(
	permission *permissionModel.Permission,
) dto.PermissionResponse {
	return dto.PermissionResponse{
		ID:          permission.ID,
		Name:        permission.Name,
		Description: permission.Description,
	}
}
