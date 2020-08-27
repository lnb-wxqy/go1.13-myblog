package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"myblog/common"
	"myblog/database"
	"myblog/model"
	"myblog/response"
	"net/http"
	"strconv"
)

//文章分类结构体

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := database.InitDB()
	//自动迁移
	db.AutoMigrate(model.Category{})

	return CategoryController{DB: db}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory model.Category
	err := json.NewDecoder(ctx.Request.Body).Decode(&requestCategory)
	if err != nil {
		response.FailedResponse(ctx, http.StatusOK, 1020, "结构体错误")
		return
	}
	if requestCategory.Name == "" {
		response.FailedResponse(ctx, http.StatusOK, 1021, "数据验证错误，分类名称不能为空")
		return
	}
	db := c.DB.Create(&requestCategory)
	if db != nil && db.Error != nil {
		response.FailedResponse(ctx, http.StatusInternalServerError, common.REGISTER_FAILED, "创建失败，"+db.Error.Error())
		return
	}
	response.SuccessResponse(ctx, gin.H{"code": "200", "requestCategory": &requestCategory})
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
	if c.DB.First(&updateCategory, categoryId).RecordNotFound() {
		response.FailedResponse(ctx, http.StatusOK, "1023", "分类不存在")
		return
	}

	// 更新分类
	requestCategory.ID = uint(categoryId)
	db := c.DB.Model(&updateCategory).Update(&requestCategory)
	if db != nil && db.Error != nil {
		response.FailedResponse(ctx, http.StatusInternalServerError, "1025", "更新失败，"+db.Error.Error())
		return
	}

	response.SuccessResponse(ctx, gin.H{"msg": "修改成功", "data": &updateCategory})
}

func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	//判断分类是否存在
	var category model.Category
	if c.DB.First(&category, categoryId).RecordNotFound() {
		response.FailedResponse(ctx, http.StatusOK, "1023", "分类不存在")
		return
	}

	response.SuccessResponse(ctx, gin.H{"data": &category})

}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if err := c.DB.Delete(model.Category{}, categoryId).Error; err != nil {
		response.FailedResponse(ctx, http.StatusOK, "1023", "删除失败")
		return
	}

	response.SuccessResponse(ctx, gin.H{"msg": "删除成功"})
}
