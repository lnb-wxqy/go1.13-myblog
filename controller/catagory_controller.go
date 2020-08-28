package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"myblog/common"
	"myblog/model"
	"myblog/repository"
	"myblog/response"
	"myblog/vo"
	"net/http"
	"strconv"
)

//文章分类结构体

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	rep repository.CategortRepository
}

func NewCategoryController() ICategoryController {

	rep := repository.NewCategoryRepository()
	//自动迁移
	rep.DB.AutoMigrate(model.Category{})

	return CategoryController{rep: rep}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CategoryRequest

	//err := json.NewDecoder(ctx.Request.Body).Decode(&requestCategory)
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.FailedResponse(ctx, http.StatusOK, 1020, err.Error())
		return
	}

	category, err := c.rep.Create(requestCategory.Name)
	if err != nil {
		response.FailedResponse(ctx, http.StatusInternalServerError, common.REGISTER_FAILED, "创建失败，"+err.Error())
		return
	}
	response.SuccessResponse(ctx, gin.H{"code": "200", "requestCategory": category})
}

func (c CategoryController) Update(ctx *gin.Context) {
	// 获取body中的参数
	var requestCategory model.Category
	err := json.NewDecoder(ctx.Request.Body).Decode(&requestCategory)
	if err != nil {
		response.FailedResponse(ctx, http.StatusOK, "1020", "结构体错误")
		return
	}

	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	//判断分类是否存在
	var updateCategory model.Category
	if c.rep.DB.First(&updateCategory, categoryId).RecordNotFound() {
		response.FailedResponse(ctx, http.StatusOK, "1023", "分类不存在")
		return
	}

	// 更新分类
	category, err := c.rep.Update(requestCategory, uint(categoryId))
	if err != nil {
		response.FailedResponse(ctx, http.StatusInternalServerError, "1025", "更新失败，"+err.Error())
		return
	}

	response.SuccessResponse(ctx, gin.H{"msg": "修改成功", "data": category})
}

func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	//判断分类是否存在
	category, err := c.rep.Select(uint(categoryId))

	if err != nil {
		response.FailedResponse(ctx, http.StatusOK, "1023", "分类不存在")
		return
	}

	response.SuccessResponse(ctx, gin.H{"data": category})

}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if err := c.rep.DeleteById(categoryId); err != nil {
		response.FailedResponse(ctx, http.StatusOK, "1023", "删除失败")
		return
	}

	response.SuccessResponse(ctx, gin.H{"msg": "删除成功"})
}
