package request

import "github.com/gin-gonic/gin"

func GetRequestID(
	ctx *gin.Context,
) string {

	value, ok := ctx.Get(RequestIDKey)
	if !ok {
		return ""
	}

	requestID, _ := value.(string)

	return requestID
}
