package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func JSON(
	ctx *gin.Context,
	statusCode int,
	success bool,
	message string,
	data any,
	err any,
) {
	ctx.JSON(statusCode, APIResponse{
		Success: success,
		Message: message,
		Data:    data,
		Error:   err,
	})
}

// Success Responses

func Success(ctx *gin.Context, message string, data any) {
	JSON(ctx, http.StatusOK, true, message, data, nil)
}

func Created(ctx *gin.Context, message string, data any) {
	JSON(ctx, http.StatusCreated, true, message, data, nil)
}

// Error Responses

func BadRequest(ctx *gin.Context, message string, err any) {
	JSON(ctx, http.StatusBadRequest, false, message, nil, err)
}

func NotFound(ctx *gin.Context, message string) {
	JSON(ctx, http.StatusNotFound, false, message, nil, nil)
}

func Conflict(ctx *gin.Context, message string) {
	JSON(ctx, http.StatusConflict, false, message, nil, nil)
}

func InternalServerError(ctx *gin.Context) {
	JSON(ctx, http.StatusInternalServerError, false, "Internal server error", nil, nil)
}
