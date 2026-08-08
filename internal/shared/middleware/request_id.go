package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/request"
	"go.uber.org/zap"
)

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		requestID := uuid.NewString()

		ctx.Set(request.RequestIDKey, requestID)

		ctx.Header("X-Request-ID", requestID)

		// TEMP DEBUG
		zap.L().Info(
			"Request ID created",
			zap.String("request_id", requestID),
		)

		ctx.Next()
	}
}
