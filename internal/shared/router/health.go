package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/request"
)

func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":     "UP",
		"service":    "AIHub Backend",
		"version":    "v1",
		"request_id": request.GetRequestID(ctx),
	})
}
