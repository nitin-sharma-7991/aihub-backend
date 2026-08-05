package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/pagination"
)

// SuccessWithMeta returns a successful response with pagination metadata.
func SuccessWithMeta(
	ctx *gin.Context,
	message string,
	data any,
	meta pagination.Meta,
) {

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": message,
		"data":    data,
		"meta":    meta,
	})
}
