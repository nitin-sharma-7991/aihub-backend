package middleware

import "github.com/gin-gonic/gin"

const (
	ContextUserID = "userID"
	ContextRole   = "role"
)

func SetUserID(ctx *gin.Context, userID uint) {
	ctx.Set(ContextUserID, userID)
}

func GetUserID(ctx *gin.Context) (uint, bool) {
	value, exists := ctx.Get(ContextUserID)
	if !exists {
		return 0, false
	}

	userID, ok := value.(uint)
	return userID, ok
}

func SetRole(ctx *gin.Context, role string) {
	ctx.Set(ContextRole, role)
}

func GetRole(ctx *gin.Context) (string, bool) {
	value, exists := ctx.Get(ContextRole)
	if !exists {
		return "", false
	}

	role, ok := value.(string)
	return role, ok
}
