package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpCode int, code, msg string) {
	ctx.JSON(httpCode, gin.H{"code": code, "msg": msg})
}

func SuccessResponse(ctx *gin.Context, msg string) {
	Response(ctx, http.StatusOK, "200", msg)
}

func FailedResponse(ctx *gin.Context, httpCode int, code, msg string) {
	Response(ctx, httpCode, code, msg)
}
