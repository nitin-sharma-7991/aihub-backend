package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/apperrors"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/response"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/security"
)

// Auth authenticates incoming requests by validating the JWT
// access token and storing authenticated user information
// in the Gin context for downstream handlers.

func Auth(secret string) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			response.Unauthorized(
				ctx,
				apperrors.ErrMissingToken.Error(),
			)
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Unauthorized(
				ctx,
				apperrors.ErrInvalidToken.Error(),
			)
			ctx.Abort()
			return
		}

		token := parts[1]

		claims, err := security.ParseJWT(
			token,
			secret,
		)

		if err != nil {
			response.Unauthorized(
				ctx,
				apperrors.ErrInvalidToken.Error(),
			)
			ctx.Abort()
			return
		}

		// Store authenticated user information.
		SetUserID(ctx, claims.UserID)
		SetRole(ctx, claims.Role)

		ctx.Next()
	}
}
