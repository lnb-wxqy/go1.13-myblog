package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"myblog/database"
	"myblog/model"
	"myblog/util"
	"net/http"
)

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
		//如果用户名为空，给一个随机的10为字符
		name = util.RandomString(10)
	}

	//密码
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422",
			"msg":  "密码不能少于6位",
		})
		return
	}

	//手机号
	if len(telephone) == 0 || len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422",
			"msg":  "手机号格式不正确",
		})
		return
	}

	//判断用户(手机号)是否存在
	if isTelePhoneExist(db, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422",
			"msg":  "用户已存在",
		})
		return
	}

	//创建用户
	newUser := &model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}

	//入库
	db.Create(newUser)

	//注册成功
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})

}

func isTelePhoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}
