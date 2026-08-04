package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/membership/dto"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/membership/service"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/apperrors"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/response"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/validation"
)

type MembershipHandler struct {
	membershipService service.MembershipService
}

func NewMembershipHandler(
	service service.MembershipService,
) *MembershipHandler {

	return &MembershipHandler{
		membershipService: service,
	}
}

// POST /memberships
func (h *MembershipHandler) Create(ctx *gin.Context) {

	var req dto.CreateMembershipRequest

	if err := validation.BindJSON(ctx, &req); err != nil {
		return
	}

	membership, err := h.membershipService.Create(
		ctx.Request.Context(),
		req,
	)

	if err != nil {

		if errors.Is(err, apperrors.ErrMembershipAlreadyExists) {
			response.Conflict(ctx, err.Error())
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Created(
		ctx,
		"Membership created successfully",
		membership,
	)
}

// GET /memberships
func (h *MembershipHandler) GetAll(ctx *gin.Context) {

	memberships, err := h.membershipService.FindAll(
		ctx.Request.Context(),
	)

	if err != nil {
		response.InternalServerError(ctx)
		return
	}

	response.Success(
		ctx,
		"Memberships fetched successfully",
		memberships,
	)
}

// GET /memberships/:id
func (h *MembershipHandler) GetByID(ctx *gin.Context) {

	id, err := strconv.ParseUint(
		ctx.Param("id"),
		10,
		64,
	)

	if err != nil {
		response.BadRequest(
			ctx,
			"Invalid membership id",
			nil,
		)
		return
	}

	membership, err := h.membershipService.FindByID(
		ctx.Request.Context(),
		uint(id),
	)

	if err != nil {

		if errors.Is(err, apperrors.ErrMembershipNotFound) {
			response.NotFound(ctx, err.Error())
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Success(
		ctx,
		"Membership fetched successfully",
		membership,
	)
}

// PUT /memberships/:id
func (h *MembershipHandler) Update(ctx *gin.Context) {

	id, err := strconv.ParseUint(
		ctx.Param("id"),
		10,
		64,
	)

	if err != nil {
		response.BadRequest(
			ctx,
			"Invalid membership id",
			nil,
		)
		return
	}

	var req dto.UpdateMembershipRequest

	if err := validation.BindJSON(ctx, &req); err != nil {
		return
	}

	membership, err := h.membershipService.Update(
		ctx.Request.Context(),
		uint(id),
		req,
	)

	if err != nil {

		if errors.Is(err, apperrors.ErrMembershipNotFound) {
			response.NotFound(ctx, err.Error())
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Success(
		ctx,
		"Membership updated successfully",
		membership,
	)
}

// DELETE /memberships/:id
func (h *MembershipHandler) Delete(ctx *gin.Context) {

	id, err := strconv.ParseUint(
		ctx.Param("id"),
		10,
		64,
	)

	if err != nil {
		response.BadRequest(
			ctx,
			"Invalid membership id",
			nil,
		)
		return
	}

	err = h.membershipService.Delete(
		ctx.Request.Context(),
		uint(id),
	)

	if err != nil {

		if errors.Is(err, apperrors.ErrMembershipNotFound) {
			response.NotFound(ctx, err.Error())
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Success(
		ctx,
		"Membership deleted successfully",
		nil,
	)
}
