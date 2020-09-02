package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//前端请求端口是8080
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", viper.GetString("auth.Access-Control-Allow-Origin"))
		ctx.Writer.Header().Set("Access-Control-Max-Age", viper.GetString("auth.Access-Control-Max-Age")) //过期时间
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", viper.GetString("auth.Access-Control-Allow-Methods"))
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", viper.GetString("auth.Access-Control-Allow-Headers"))
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", viper.GetString("auth.Access-Control-Allow-Credentials"))

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusOK)
		} else {
			ctx.Next()
		}
	}
}
