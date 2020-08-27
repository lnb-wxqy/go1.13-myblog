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
	// 跨域拦截器
	r.Use(middleware.CORSMiddleware())
	//用户相关
	group := r.Group(common.USER_GROUP)
	group.Handle(http.MethodPost, common.REGISTER, controller.Register)
	group.Handle(http.MethodPost, common.LOGIN, controller.Login)
	group.Handle(http.MethodGet, common.INFO, middleware.AuthMiddleWare(), controller.Info)

	// 文章相关
	categoryGroup := r.Group(common.CATEGORY_GROUP)
	categoryController := controller.NewCategoryController()
	categoryGroup.Handle(http.MethodPost, common.CREATE, categoryController.Create)
	categoryGroup.Handle(http.MethodPut, common.UPDATE+"/:id", categoryController.Update)
	categoryGroup.Handle(http.MethodGet, common.SHOW+"/:id", categoryController.Show)
	categoryGroup.Handle(http.MethodDelete, common.DELETE+"/:id", categoryController.Delete)

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run(":8080"))
}
