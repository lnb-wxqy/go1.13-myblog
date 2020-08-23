package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"myblog/common"
	"myblog/controller"
	"myblog/middleware"
	"net/http"
)

func StartProxy() {
	r := gin.Default()
	group := r.Group(common.GROUP)
	group.Handle(http.MethodPost, common.REGISTER, controller.Register)
	group.Handle(http.MethodPost, common.LOGIN, controller.Login)
	group.Handle(http.MethodPost, common.INFO, middleware.AuthMiddleWare(), controller.Info)

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run(":8080"))
}
