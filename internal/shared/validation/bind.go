package validation

import (
	"github.com/gin-gonic/gin"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/response"
)

// BindJSON binds and validates a JSON request body.
func BindJSON(
	ctx *gin.Context,
	req any,
) error {

	if err := ctx.ShouldBindJSON(req); err != nil {

		response.BadRequest(
			ctx,
			"Validation failed",
			FormatErrors(err),
		)

		return err
	}

	return nil
}
