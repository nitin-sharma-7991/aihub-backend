package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/organization/dto"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/organization/service"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/apperrors"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/pagination"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/response"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/validation"
)

type OrganizationHandler struct {
	service service.OrganizationService
}

func NewOrganizationHandler(service service.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{
		service: service,
	}
}

// POST /organization
func (h *OrganizationHandler) Create(ctx *gin.Context) {

	var req dto.CreateOrganizationRequest

	if err := validation.BindJSON(ctx, &req); err != nil {
		return
	}

	org, err := h.service.Create(ctx.Request.Context(), req)
	if err != nil {

		if errors.Is(err, apperrors.ErrOrganizationAlreadyExists) {
			response.Conflict(ctx, err.Error())
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Created(ctx, "Organization created successfully", org)
}

func (h *OrganizationHandler) GetAll(ctx *gin.Context) {

	var req pagination.Request

	ctx.ShouldBindQuery(&req)

	orgs, meta, err := h.service.FindAll(
		ctx.Request.Context(),
		req,
	)

	if err != nil {
		response.InternalServerError(ctx)
		return
	}

	response.SuccessWithMeta(
		ctx,
		"Organizations fetched successfully",
		orgs,
		meta,
	)
}

// GET /organizations/:organization_id
func (h *OrganizationHandler) GetByID(ctx *gin.Context) {

	id, err := strconv.ParseUint(ctx.Param("organization_id"), 10, 64)
	if err != nil {
		response.BadRequest(ctx, "Invalid organization id", nil)
		return
	}

	org, err := h.service.FindByID(ctx.Request.Context(), uint(id))
	if err != nil {

		if errors.Is(err, apperrors.ErrOrganizationNotFound) {
			response.NotFound(ctx, err.Error())
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Success(ctx, "Organization fetched successfully", org)
}

// PUT /organizations/:organization_id
func (h *OrganizationHandler) Update(ctx *gin.Context) {

	id, err := strconv.ParseUint(ctx.Param("organization_id"), 10, 64)
	if err != nil {
		response.BadRequest(ctx, "Invalid organization id", nil)
		return
	}

	var req dto.UpdateOrganizationRequest

	if err := validation.BindJSON(ctx, &req); err != nil {
		return
	}

	org, err := h.service.Update(
		ctx.Request.Context(),
		uint(id),
		req,
	)

	if err != nil {

		if errors.Is(err, apperrors.ErrOrganizationNotFound) {
			response.NotFound(ctx, err.Error())
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Success(ctx, "Organization updated successfully", org)
}

// DELETE /organizations/:id
func (h *OrganizationHandler) Delete(ctx *gin.Context) {

	id, err := strconv.ParseUint(ctx.Param("organization_id"), 10, 64)
	if err != nil {
		response.BadRequest(ctx, "Invalid organization id", nil)
		return
	}

	err = h.service.Delete(ctx.Request.Context(), uint(id))
	if err != nil {

		if errors.Is(err, apperrors.ErrOrganizationNotFound) {
			response.NotFound(ctx, err.Error())
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Success(ctx, "Organization deleted successfully", nil)
}
