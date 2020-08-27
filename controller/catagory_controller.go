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


}

func (c CategoryController) Show(ctx *gin.Context) {
	id :=ctx.Param("id")
	if category, b := getCategory(c.DB, id); b {
		response.SuccessResponse(ctx, gin.H{"code": "200", "data": category})
		return
	}
	return
}

func (c CategoryController) Delete(ctx *gin.Context) {
	panic("implement me")
}

func getCategory(db *gorm.DB, id string) (*model.Category, bool) {
	var category model.Category
	db.Where("name = ?", name).First(&category)
	if category.ID == 0 {
		return nil, false
	}

	return &category, true
}
