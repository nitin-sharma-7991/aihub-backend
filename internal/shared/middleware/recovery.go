package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Recovery(log *zap.Logger) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		defer func() {

			if err := recover(); err != nil {

				log.Error(
					"Panic recovered",
					zap.Any("error", err),
					zap.ByteString("stack", debug.Stack()),
				)

				ctx.AbortWithStatusJSON(
					http.StatusInternalServerError,
					gin.H{
						"success": false,
						"message": "Internal Server Error",
					},
				)
			}

		}()

		ctx.Next()
	}
}
