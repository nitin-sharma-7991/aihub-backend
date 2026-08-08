package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/request"
	"go.uber.org/zap"
)

func Logger(log *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		start := time.Now()

		ctx.Next()

		latency := time.Since(start)

		requestID := request.GetRequestID(ctx)

		log.Info(
			"HTTP Request",
			zap.String("request_id", requestID),
			zap.String("method", ctx.Request.Method),
			zap.String("path", ctx.Request.URL.Path),
			zap.Int("status", ctx.Writer.Status()),
			zap.Duration("latency", latency),
			zap.String("client_ip", ctx.ClientIP()),
			zap.String("user_agent", ctx.Request.UserAgent()),
		)
	}
}
