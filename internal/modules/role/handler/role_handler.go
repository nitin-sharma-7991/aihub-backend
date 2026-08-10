package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/role/dto"
	roleModel "github.com/nitin-sharma-7991/aihub-backend/internal/modules/role/model"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/role/service"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/apperrors"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/pagination"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/response"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/validation"
)

type RoleHandler struct {
	roleService service.RoleService
}

func NewRoleHandler(
	roleService service.RoleService,
) *RoleHandler {
	return &RoleHandler{
		roleService: roleService,
	}
}

// POST /roles
func (h *RoleHandler) Create(ctx *gin.Context) {

	var req dto.CreateRoleRequest

	if err := validation.BindJSON(ctx, &req); err != nil {
		return
	}

	role, err := h.roleService.Create(
		ctx.Request.Context(),
		req.Name,
		req.Description,
	)

	if err != nil {

		if errors.Is(err, apperrors.ErrInvalidRoleName) {
			response.BadRequest(
				ctx,
				err.Error(),
				nil,
			)
			return
		}

		if errors.Is(err, apperrors.ErrRoleAlreadyExists) {
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
		"Role created successfully",
		toRoleResponse(role),
	)
}

// GET /roles
func (h *RoleHandler) GetAll(ctx *gin.Context) {

	var req pagination.Request

	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.BadRequest(
			ctx,
			"Invalid pagination parameters",
			nil,
		)
		return
	}

	roles, meta, err := h.roleService.FindAll(
		ctx.Request.Context(),
		req,
	)

	if err != nil {
		response.InternalServerError(ctx)
		return
	}

	responses := make([]dto.RoleResponse, 0, len(roles))

	for _, role := range roles {
		responses = append(
			responses,
			toRoleResponse(&role),
		)
	}

	response.SuccessWithMeta(
		ctx,
		"Roles fetched successfully",
		responses,
		meta,
	)
}

// GET /roles/:id
func (h *RoleHandler) GetByID(ctx *gin.Context) {

	id, err := strconv.ParseUint(
		ctx.Param("id"),
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

	role, err := h.roleService.FindByID(
		ctx.Request.Context(),
		uint(id),
	)

	if err != nil {

		if errors.Is(err, apperrors.ErrRoleNotFound) {
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
		"Role fetched successfully",
		toRoleResponse(role),
	)
}

// PUT /roles/:id
func (h *RoleHandler) Update(ctx *gin.Context) {

	id, err := strconv.ParseUint(
		ctx.Param("id"),
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

	var req dto.UpdateRoleRequest

	if err := validation.BindJSON(ctx, &req); err != nil {
		return
	}

	role, err := h.roleService.Update(
		ctx.Request.Context(),
		uint(id),
		req.Name,
		req.Description,
	)

	if err != nil {

		if errors.Is(err, apperrors.ErrRoleNotFound) {
			response.NotFound(
				ctx,
				err.Error(),
			)
			return
		}

		if errors.Is(err, apperrors.ErrInvalidRoleName) {
			response.BadRequest(
				ctx,
				err.Error(),
				nil,
			)
			return
		}

		if errors.Is(err, apperrors.ErrRoleAlreadyExists) {
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
		"Role updated successfully",
		toRoleResponse(role),
	)
}

// DELETE /roles/:id
func (h *RoleHandler) Delete(ctx *gin.Context) {

	id, err := strconv.ParseUint(
		ctx.Param("id"),
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

	err = h.roleService.Delete(
		ctx.Request.Context(),
		uint(id),
	)

	if err != nil {

		if errors.Is(err, apperrors.ErrRoleNotFound) {
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
		"Role deleted successfully",
		nil,
	)
}

func toRoleResponse(
	role *roleModel.Role,
) dto.RoleResponse {
	return dto.RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
	}
}
