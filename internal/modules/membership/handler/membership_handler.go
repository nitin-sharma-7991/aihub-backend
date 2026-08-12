package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/membership/dto"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/membership/service"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/apperrors"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/pagination"
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

// ================================================================
// CREATE
// POST /organizations/:organization_id/memberships
// ================================================================

func (h *MembershipHandler) Create(
	ctx *gin.Context,
) {

	organizationID, ok := getOrganizationID(ctx)
	if !ok {
		response.BadRequest(
			ctx,
			"Invalid organization context",
			nil,
		)
		return
	}

	var req dto.CreateMembershipRequest

	if err := validation.BindJSON(
		ctx,
		&req,
	); err != nil {
		return
	}

	membership, err := h.membershipService.Create(
		ctx.Request.Context(),
		organizationID,
		req,
	)

	if err != nil {
		if errors.Is(
			err,
			apperrors.ErrMembershipAlreadyExists,
		) {
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
		"Membership created successfully",
		membership,
	)
}

// ================================================================
// GET ALL GLOBAL
// GET /memberships
// ================================================================

func (h *MembershipHandler) GetAll(
	ctx *gin.Context,
) {

	var req pagination.Request

	ctx.ShouldBindQuery(&req)

	memberships, meta, err :=
		h.membershipService.FindAll(
			ctx.Request.Context(),
			req,
		)

	if err != nil {
		response.InternalServerError(ctx)
		return
	}

	response.SuccessWithMeta(
		ctx,
		"Memberships fetched successfully",
		memberships,
		meta,
	)
}

// ================================================================
// GET ALL BY ORGANIZATION
// GET /organizations/:organization_id/memberships
// ================================================================

func (h *MembershipHandler) GetAllByOrganization(
	ctx *gin.Context,
) {

	orgID, ok := getOrganizationID(ctx)
	if !ok {
		response.BadRequest(
			ctx,
			"Invalid organization context",
			nil,
		)
		return
	}

	var req pagination.Request

	if err := ctx.ShouldBindQuery(
		&req,
	); err != nil {
		response.BadRequest(
			ctx,
			"Invalid pagination parameters",
			err.Error(),
		)
		return
	}

	memberships, meta, err :=
		h.membershipService.FindAllByOrganizationID(
			ctx.Request.Context(),
			orgID,
			req,
		)

	if err != nil {
		response.InternalServerError(ctx)
		return
	}

	response.SuccessWithMeta(
		ctx,
		"Organization memberships fetched successfully",
		memberships,
		meta,
	)
}

// ================================================================
// GET BY ID GLOBAL
// GET /memberships/:id
// ================================================================

func (h *MembershipHandler) GetByID(
	ctx *gin.Context,
) {

	id, err := parseMembershipID(ctx)
	if err != nil {
		return
	}

	membership, err := h.membershipService.FindByID(
		ctx.Request.Context(),
		id,
	)

	if err != nil {
		if errors.Is(
			err,
			apperrors.ErrMembershipNotFound,
		) {
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
		"Membership fetched successfully",
		membership,
	)
}

// ================================================================
// GET BY ID + ORGANIZATION
// GET /organizations/:organization_id/memberships/:id
// ================================================================

func (h *MembershipHandler) GetByIDAndOrganization(
	ctx *gin.Context,
) {

	organizationID, ok := getOrganizationID(ctx)
	if !ok {
		response.BadRequest(
			ctx,
			"Invalid organization context",
			nil,
		)
		return
	}

	id, err := parseMembershipID(ctx)
	if err != nil {
		return
	}

	membership, err :=
		h.membershipService.FindByIDAndOrganizationID(
			ctx.Request.Context(),
			id,
			organizationID,
		)

	if err != nil {
		if errors.Is(
			err,
			apperrors.ErrMembershipNotFound,
		) {
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
		"Membership fetched successfully",
		membership,
	)
}

// ================================================================
// UPDATE GLOBAL
// PUT /memberships/:id
// ================================================================

func (h *MembershipHandler) Update(
	ctx *gin.Context,
) {

	id, err := parseMembershipID(ctx)
	if err != nil {
		return
	}

	var req dto.UpdateMembershipRequest

	if err := validation.BindJSON(
		ctx,
		&req,
	); err != nil {
		return
	}

	membership, err := h.membershipService.Update(
		ctx.Request.Context(),
		id,
		req,
	)

	if err != nil {
		if errors.Is(
			err,
			apperrors.ErrMembershipNotFound,
		) {
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
		"Membership updated successfully",
		membership,
	)
}

// ================================================================
// UPDATE BY ORGANIZATION
// PUT /organizations/:organization_id/memberships/:id
// ================================================================

func (h *MembershipHandler) UpdateByOrganization(
	ctx *gin.Context,
) {

	organizationID, ok := getOrganizationID(ctx)
	if !ok {
		response.BadRequest(
			ctx,
			"Invalid organization context",
			nil,
		)
		return
	}

	id, err := parseMembershipID(ctx)
	if err != nil {
		return
	}

	var req dto.UpdateMembershipRequest

	if err := validation.BindJSON(
		ctx,
		&req,
	); err != nil {
		return
	}

	membership, err :=
		h.membershipService.UpdateByOrganizationID(
			ctx.Request.Context(),
			id,
			organizationID,
			req,
		)

	if err != nil {
		if errors.Is(
			err,
			apperrors.ErrMembershipNotFound,
		) {
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
		"Membership updated successfully",
		membership,
	)
}

// ================================================================
// DELETE GLOBAL
// DELETE /memberships/:id
// ================================================================

func (h *MembershipHandler) Delete(
	ctx *gin.Context,
) {

	id, err := parseMembershipID(ctx)
	if err != nil {
		return
	}

	err = h.membershipService.Delete(
		ctx.Request.Context(),
		id,
	)

	if err != nil {
		if errors.Is(
			err,
			apperrors.ErrMembershipNotFound,
		) {
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
		"Membership deleted successfully",
		nil,
	)
}

// ================================================================
// DELETE BY ORGANIZATION
// DELETE /organizations/:organization_id/memberships/:id
// ================================================================

func (h *MembershipHandler) DeleteByOrganization(
	ctx *gin.Context,
) {

	organizationID, ok := getOrganizationID(ctx)
	if !ok {
		response.BadRequest(
			ctx,
			"Invalid organization context",
			nil,
		)
		return
	}

	id, err := parseMembershipID(ctx)
	if err != nil {
		return
	}

	err = h.membershipService.DeleteByOrganizationID(
		ctx.Request.Context(),
		id,
		organizationID,
	)

	if err != nil {
		if errors.Is(
			err,
			apperrors.ErrMembershipNotFound,
		) {
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
		"Membership deleted successfully",
		nil,
	)
}

// ================================================================
// HELPERS
// ================================================================

func getOrganizationID(
	ctx *gin.Context,
) (uint, bool) {

	value, exists := ctx.Get(
		"organization_id",
	)
	if !exists {
		return 0, false
	}

	organizationID, ok := value.(uint)
	if !ok || organizationID == 0 {
		return 0, false
	}

	return organizationID, true
}

func parseMembershipID(
	ctx *gin.Context,
) (uint, error) {

	id, err := strconv.ParseUint(
		ctx.Param("id"),
		10,
		64,
	)

	if err != nil || id == 0 {
		response.BadRequest(
			ctx,
			"Invalid membership id",
			nil,
		)

		return 0, errors.New("invalid membership id")
	}

	return uint(id), nil
}
