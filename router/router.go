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
	//用户相关接口
	group := r.Group(common.USER_GROUP)
	group.Handle(http.MethodPost, common.REGISTER, controller.Register)
	group.Handle(http.MethodPost, common.LOGIN, controller.Login)
	group.Handle(http.MethodGet, common.INFO, middleware.AuthMiddleWare(), controller.Info)

	// 文章分类接口
	categoryGroup := r.Group(common.CATEGORY_GROUP)
	categoryController := controller.NewCategoryController()
	categoryGroup.Handle(http.MethodPost, common.CREATE, categoryController.Create)
	categoryGroup.Handle(http.MethodPut, common.UPDATE+"/:id", categoryController.Update)
	categoryGroup.Handle(http.MethodGet, common.SHOW+"/:id", categoryController.Show)
	categoryGroup.Handle(http.MethodDelete, common.DELETE+"/:id", categoryController.Delete)

	// 文章接口
	articleGroup := r.Group(common.ARTICLE_GROUP)
	articleGroup.Use(middleware.AuthMiddleWare()) //中间件，获取ctx中的用户信息
	articleController := controller.NewArticleController()
	articleGroup.Handle(http.MethodPost, common.CREATE, articleController.Create)
	articleGroup.Handle(http.MethodPut, common.UPDATE+"/:id", articleController.Update)
	articleGroup.Handle(http.MethodGet, common.SHOW+"/:id", articleController.Show)
	articleGroup.Handle(http.MethodDelete, common.DELETE+"/:id", articleController.Delete)
	articleGroup.Handle(http.MethodPost, common.PAGELIST, articleController.PageList)

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run(":8080"))
}
