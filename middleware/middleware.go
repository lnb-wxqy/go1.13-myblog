package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"myblog/common"
	"myblog/database"
	"myblog/model"
	"myblog/response"
	"net/http"
	"strings"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		header := ctx.GetHeader("Authorization")
		//Bearer auth2.0规定Authorization必须以Bearer开头
		//validate header format
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": "401", "msg": "权限不足"})
			ctx.Abort() //丢弃该请求
			return
		}

		//截取token的有效部分，因为 "Bearer "占了7位，故从第7位开始截取
		tokenString := header[7:]
		//解析token
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			if validationError, ok := err.(*jwt.ValidationError); ok {
				response.Response(ctx, http.StatusUnauthorized, common.TOKEN_IS_INVALID, validationError.Inner.Error())
			} else {
				response.Response(ctx, http.StatusInternalServerError, common.STATUS_INTERNAL_SERVER_ERROR, common.StatusText[common.STATUS_INTERNAL_SERVER_ERROR])
			}
			ctx.Abort() //丢弃该请求
			return
		}

		//验证通过，从claims中获取用户id
		id := claims.UserId
		db := database.InitDB()
		defer db.Close()

		//验证用户是否存在
		//不存在，返回权限不足
		var user model.User
		db.First(&user, id)
		if user.ID == 0 {
			response.Response(ctx, http.StatusUnprocessableEntity, common.USER_IS_NOT_EXIST, common.StatusText[common.USER_IS_NOT_EXIST])
			ctx.Abort() //丢弃该请求
			return
		}

		//存在，将用户信息写入上下文
		ctx.Set("user", &user)

		// Next should be used only inside middleware.
		// It executes the pending handlers in the chain inside the calling handler.
		ctx.Next()

	}
}
