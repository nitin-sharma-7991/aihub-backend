package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/authorization"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/response"
)

func RequirePermission(
	authorizationService authorization.Service,
	permission string,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Get authenticated user ID from Auth middleware.
		userIDValue, exists := ctx.Get("user_id")
		if !exists {
			response.Unauthorized(
				ctx,
				"Unauthorized",
			)
			ctx.Abort()
			return
		}

		userID, ok := userIDValue.(uint)
		if !ok {
			response.Unauthorized(
				ctx,
				"Invalid authentication context",
			)
			ctx.Abort()
			return
		}

		// Organization context must be set before permission middleware.
		organizationIDValue, exists := ctx.Get("organization_id")
		if !exists {
			response.BadRequest(
				ctx,
				"Organization context is required",
				nil,
			)
			ctx.Abort()
			return
		}

		organizationID, ok := organizationIDValue.(uint)
		if !ok {
			response.BadRequest(
				ctx,
				"Invalid organization context",
				nil,
			)
			ctx.Abort()
			return
		}

		allowed, err := authorizationService.Can(
			ctx.Request.Context(),
			userID,
			organizationID,
			permission,
		)

		if err != nil {
			response.InternalServerError(ctx)
			ctx.Abort()
			return
		}

		if !allowed {
			response.Forbidden(
				ctx,
				"Permission denied",
			)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
