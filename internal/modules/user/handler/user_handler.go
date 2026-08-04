package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/user/dto"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/user/service"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/apperrors"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/response"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/validation"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// POST /users
func (h *UserHandler) Create(ctx *gin.Context) {

	var req dto.CreateUserRequest

	if err := validation.BindJSON(ctx, &req); err != nil {
		return
	}

	user, err := h.userService.Create(ctx.Request.Context(), req)
	if err != nil {

		if errors.Is(err, apperrors.ErrEmailAlreadyExists) {
			response.Conflict(ctx, err.Error())
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Created(ctx, "User created successfully", user)
}

// GET /users/:id
func (h *UserHandler) GetByID(ctx *gin.Context) {

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(ctx, "Invalid user id", nil)
		return
	}

	user, err := h.userService.GetByID(ctx.Request.Context(), uint(id))
	if err != nil {

		if errors.Is(err, apperrors.ErrUserNotFound) {
			response.NotFound(ctx, err.Error())
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Success(ctx, "User fetched successfully", user)
}

// PUT /users/:id
func (h *UserHandler) Update(ctx *gin.Context) {

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(ctx, "Invalid user id", nil)
		return
	}

	var req dto.UpdateUserRequest

	if err := validation.BindJSON(ctx, &req); err != nil {
		return
	}

	user, err := h.userService.Update(
		ctx.Request.Context(),
		uint(id),
		req,
	)

	if err != nil {

		if errors.Is(err, apperrors.ErrUserNotFound) {
			response.NotFound(ctx, err.Error())
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Success(ctx, "User updated successfully", user)
}

// DELETE /users/:id
func (h *UserHandler) Delete(ctx *gin.Context) {

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(ctx, "Invalid user id", nil)
		return
	}

	err = h.userService.Delete(ctx.Request.Context(), uint(id))
	if err != nil {

		if errors.Is(err, apperrors.ErrUserNotFound) {
			response.NotFound(ctx, err.Error())
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Success(ctx, "User deleted successfully", nil)
}
