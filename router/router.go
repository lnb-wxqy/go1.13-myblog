package router

import (
	"github.com/gin-gonic/gin"
	"myblog/common"
	"myblog/controller"
	"net/http"
)

func StartProxy() {
	r := gin.Default()
	group := r.Group(common.GROUP)
	group.Handle(http.MethodPost, common.REGISTER, controller.Register)
	r.Run(":8080")
}
