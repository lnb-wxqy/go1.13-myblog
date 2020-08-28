package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"myblog/database"
	"myblog/model"
	"myblog/response"
	"myblog/vo"
	"net/http"
	"strconv"
)

type IArticleController interface {
	RestController
	PageList(ctx *gin.Context)
}

type ArticleController struct {
	DB *gorm.DB
}

func NewArticleController() IArticleController {
	db := database.InitDB()
	db.AutoMigrate(model.Article{})
	return ArticleController{DB: db}
}

func (a ArticleController) Create(ctx *gin.Context) {
	// 获取参数,并验证
	var articleRequest vo.ArticleRequest
	if err := ctx.ShouldBind(&articleRequest); err != nil {
		response.FailedResponse(ctx, http.StatusOK, 1020, err.Error())
		return
	}

	// 从上下文中获取用户信息
	user, _ := ctx.Get("user")

	// 创建文章
	article := model.Article{
		UserId:     user.(*model.User).ID,
		CategoryId: articleRequest.CategoryId,
		Title:      articleRequest.Title,
		HeadImg:    articleRequest.HeadImg,
		Content:    articleRequest.Content,
	}

	if err := a.DB.Create(&article).Error; err != nil {
		response.FailedResponse(ctx, http.StatusOK, "1031", "文章创建失败"+err.Error())
		return
	}

	response.SuccessResponse(ctx, gin.H{"code": "200", "msg": "文章创建成功"})
}

func (a ArticleController) Update(ctx *gin.Context) {
	// 获取参数,并验证
	var articleRequest vo.ArticleRequest
	if err := ctx.ShouldBind(&articleRequest); err != nil {
		response.FailedResponse(ctx, http.StatusOK, "1032", err.Error())
		return
	}

	// 判断文章是否存在
	// 从path中获取文章id
	articleId := ctx.Params.ByName("id")
	var article model.Article
	if a.DB.Where("id = ?", articleId).First(&article).RecordNotFound() {
		response.FailedResponse(ctx, http.StatusOK, "1033", "文章不存在")
		return
	}

	// 从上下文中获取用户信息
	user, _ := ctx.Get("user")

	// 判断文章是否属于该用户
	userId := user.(*model.User).ID
	if userId != article.UserId {
		response.FailedResponse(ctx, http.StatusOK, "1033", "文章不属于您，请勿非法操作")
		return
	}

	//更新
	if err := a.DB.Model(&article).Update(articleRequest).Error; err != nil {
		response.FailedResponse(ctx, http.StatusOK, "1033", "文章更新失败， "+err.Error())
		return
	}

	response.SuccessResponse(ctx, gin.H{"code": 200, "data": article, "msg": "更新成功"})

}

func (a ArticleController) Show(ctx *gin.Context) {
	// 从path中获取文章id
	articleId := ctx.Params.ByName("id")
	var article = &model.Article{}
	// gorm的关联查询
	a.DB.Preload("Category").Where("id = ?", articleId).First(article)
	response.FailedResponse(ctx, http.StatusOK, "200", gin.H{
		"data": article,
	})

}

func (a ArticleController) Delete(ctx *gin.Context) {
	// 从path中获取文章id
	articleId := ctx.Params.ByName("id")

	// 判断文章是否存在
	var article model.Article
	if a.DB.Where("id = ?", articleId).First(&article).RecordNotFound() {
		response.FailedResponse(ctx, http.StatusOK, "1033", "删除失败，文章不存在")
		return
	}

	// 从上下文中获取用户信息
	user, _ := ctx.Get("user")

	// 判断文章是否属于该用户
	userId := user.(*model.User).ID
	if userId != article.UserId {
		response.FailedResponse(ctx, http.StatusOK, "1033", "文章不属于您，请勿非法操作")
		return
	}

	if err := a.DB.Delete(&article).Error; err != nil {
		response.FailedResponse(ctx, http.StatusOK, "1035", "删除文件失败，"+err.Error())
		return
	}

	response.SuccessResponse(ctx, gin.H{"msg": "删除成功", "data": &article})
}

// 分页查询
func (a ArticleController) PageList(ctx *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	// 分页
	var articles []model.Article
	if err := a.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&articles).Error; err != nil {
		response.FailedResponse(ctx, http.StatusOK, "1040", "分页查询失败, "+err.Error())
		return
	}

	// 前端渲染需要知道总数
	var total int
	if err := a.DB.Model(model.Article{}).Count(&total).Error; err != nil {
		response.FailedResponse(ctx, http.StatusOK, "1040", "查询总数失败, "+err.Error())
		return
	}

	response.SuccessResponse(ctx, gin.H{"msg": "成功", "data": articles, "total": total})
}
