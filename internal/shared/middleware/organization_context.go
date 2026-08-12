package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/response"
)

func OrganizationContext() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		organizationIDParam := ctx.Param("organization_id")

		organizationID, err := strconv.ParseUint(
			organizationIDParam,
			10,
			64,
		)

		if err != nil || organizationID == 0 {
			response.BadRequest(
				ctx,
				"Invalid organization id",
				nil,
			)
			ctx.Abort()
			return
		}

		ctx.Set(
			"organization_id",
			uint(organizationID),
		)

		ctx.Next()
	}
}
