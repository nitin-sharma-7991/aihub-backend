package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/apperrors"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/response"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/security"
)

// Auth validates JWT access tokens.
func Auth(secret string) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			response.Forbidden(
				ctx,
				apperrors.ErrMissingToken.Error(),
			)
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Forbidden(
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
			response.Forbidden(
				ctx,
				apperrors.ErrInvalidToken.Error(),
			)
			ctx.Abort()
			return
		}

		// Store authenticated user information.
		ctx.Set("userID", claims.UserID)
		ctx.Set("role", claims.Role)

		ctx.Next()
	}
}
