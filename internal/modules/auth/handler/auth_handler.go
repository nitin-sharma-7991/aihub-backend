package handler

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/auth/dto"
	"github.com/nitin-sharma-7991/aihub-backend/internal/modules/auth/service"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/apperrors"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/response"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/validation"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(ctx *gin.Context) {

	var req dto.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {

		response.BadRequest(
			ctx,
			"Validation failed",
			validation.FormatErrors(err),
		)

		return
	}

	loginResponse, err := h.authService.Login(
		ctx.Request.Context(),
		req,
	)

	if err != nil {

		if errors.Is(err, apperrors.ErrInvalidCredentials) {
			response.Forbidden(ctx, err.Error())
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Success(
		ctx,
		"Login successful",
		loginResponse,
	)
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	// TODO
}

func (h *AuthHandler) Me(ctx *gin.Context) {

	userIDValue, exists := ctx.Get("userID")
	if !exists {
		response.Forbidden(
			ctx,
			apperrors.ErrUnauthorized.Error(),
		)
		return
	}

	userID, ok := userIDValue.(uint)
	if !ok {
		response.Forbidden(
			ctx,
			apperrors.ErrUnauthorized.Error(),
		)
		return
	}

	user, err := h.authService.Me(
		ctx.Request.Context(),
		userID,
	)

	if err != nil {

		if errors.Is(err, apperrors.ErrUserNotFound) {
			response.NotFound(ctx, err.Error())
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Success(
		ctx,
		"User profile fetched successfully",
		user,
	)
}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	// TODO
}

func (h *AuthHandler) RefreshToken(ctx *gin.Context) {
	// TODO
}
