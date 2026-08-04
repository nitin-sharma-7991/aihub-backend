package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/request"
)

func RequestID() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		requestID := uuid.NewString()

		ctx.Set(
			request.RequestIDKey,
			requestID,
		)

		ctx.Writer.Header().Set(
			"X-Request-ID",
			requestID,
		)

		ctx.Next()
	}
}
