package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"myblog/common"
	"myblog/database"
	"myblog/dto"
	"myblog/model"
	"myblog/response"
	"myblog/util"
	"net/http"
)

// 注册
func Register(c *gin.Context) {
	db := database.InitDB()
	defer db.Close()
	//获取参数
	//使用map获取参数
	//var m = make(map[string]string, 0)
	//json.NewDecoder(c.Request.Body).Decode(&m)

	var user = &model.User{}
	json.NewDecoder(c.Request.Body).Decode(user)

	name := user.Name
	password := user.Password
	telephone := user.Telephone

	//数据验证
	//用户名
	if len(name) == 0 {
		response.Response(c, http.StatusUnprocessableEntity, common.USERNAME_IS_NULL, common.StatusText[common.USERNAME_IS_NULL])
		return
	}
	//判断用户是否存在
	if isUserExist(db, name) {
		response.Response(c, http.StatusUnprocessableEntity, common.USER_IS_EXIST, common.StatusText[common.USER_IS_EXIST])
		return
	}

	//密码
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, common.PASSWORD_LENGTH_LESS_SIX, common.StatusText[common.PASSWORD_LENGTH_LESS_SIX])
		return
	}

	//手机号允许为空
	if len(telephone) != 0 && len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, common.TELEPHONE_FORMAT_ERROR, common.StatusText[common.TELEPHONE_FORMAT_ERROR])
		return
	}

	//用户密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, common.GENERATE_FROM_PASSWORD, common.StatusText[common.GENERATE_FROM_PASSWORD])
		return
	}

	//创建用户
	newUser := &model.User{
		Name:      name,
		Telephone: telephone,
		Password:  util.BytesString(hashPassword),
	}

	//入库
	ret := db.Create(newUser)
	if ret != nil && ret.Error != nil {
		response.FailedResponse(c, http.StatusInternalServerError, common.REGISTER_FAILED, "注册失败，"+ret.Error.Error())
		return
	}
	//发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, common.STATUS_INTERNAL_SERVER_ERROR, common.StatusText[common.STATUS_INTERNAL_SERVER_ERROR])
		return
	}
	//注册成功
	c.JSON(http.StatusOK, gin.H{
		"code":  "200",
		"token": token,
		"msg":   "注册成功",
	})

}

// 登录
func Login(c *gin.Context) {

	////获取参数
	var user = &model.User{}
	json.NewDecoder(c.Request.Body).Decode(user)
	name := user.Name
	password := user.Password
	//telephone := c.PostForm("telephone")

	db := database.InitDB()
	defer db.Close()
	//数据验证
	//判断用户是否存在
	var newUser model.User
	db.Where("name = ?", name).First(&newUser)
	if newUser.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, common.USER_IS_NOT_EXIST, common.StatusText[common.USER_IS_NOT_EXIST])
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusUnprocessableEntity, common.PASSWORD_IS_NOT_CORRECT, common.StatusText[common.PASSWORD_IS_NOT_CORRECT])
		return
	}

	//发放token
	token, err := common.ReleaseToken(&newUser)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, common.STATUS_INTERNAL_SERVER_ERROR, common.StatusText[common.STATUS_INTERNAL_SERVER_ERROR])
		return
	}

	//登录成功
	c.JSON(http.StatusOK, gin.H{
		"code":  "200",
		"token": token,
		"msg":   "登录成功",
	})
}

func isUserExist(db *gorm.DB, name string) bool {
	var user model.User
	db.Where("name = ?", name).First(&user)
	return user.ID != 0
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": "200", "user": dto.ToUserDto(user.(*model.User))})
}
