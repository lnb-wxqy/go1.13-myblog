package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"myblog/database"
	"myblog/model"
	"myblog/util"
	"net/http"
)

// 注册
func Register(c *gin.Context) {
	db := database.InitDB()
	defer db.Close()
	//获取参数
	name := c.PostForm("name")
	password := c.PostForm("password")
	telephone := c.PostForm("telephone")

	//数据验证
	//用户名
	if len(name) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422",
			"msg":  "用户名不能为空",
		})
		return
	}
	//判断用户是否存在
	if isUserExist(db, name) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422",
			"msg":  "用户已存在",
		})
		return
	}

	//密码
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422",
			"msg":  "密码不能少于6位",
		})
		return
	}

	//手机号允许为空
	if len(telephone) != 0 && len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422",
			"msg":  "手机号格式不正确",
		})
		return
	}

	//用户密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Println("hashPassword length: ", len(hashPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": "500",
			"msg":  "加密错误",
		})
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": "500",
			"msg":  "注册失败",
			"err:": ret.Error.Error(),
		})
		return
	}

	//注册成功
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "注册成功",
	})

}

// 登录
func Login(c *gin.Context) {

	//获取参数
	name := c.PostForm("name")
	password := c.PostForm("password")
	//telephone := c.PostForm("telephone")

	db := database.InitDB()
	defer db.Close()
	//数据验证
	//判断用户是否存在
	var user model.User
	db.Where("name = ?", name).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422",
			"msg":  "用户不存在",
		})
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422",
			"msg":  "密码不正确",
		})
		return
	}

	//发放token
	token := "11"

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
