package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"type:varchar(50);not null"`
}

func main() {

	db := InitDB()
	defer db.Close()

	server := gin.Default()
	server.POST("/api/auth/register", func(c *gin.Context) {
		//获取参数
		name := c.PostForm("name")
		password := c.PostForm("password")
		telephone := c.PostForm("telephone")

		//数据验证
		//用户名
		if len(name) == 0 {
			//如果用户名为空，给一个随机的10为字符
			name = RandomString(10)
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
		newUser :=&User{
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

	})

	server.Run(":8080")
}

func RandomString(n int) string {

	letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

//数据库 gorm
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "39.106.0.106"
	port := "3306"
	database := "myblog"
	username := "root"
	password := "Umissi@213"
	charset := "utf8"
	//"mysql":      "mysql",    // root:Netposa123@tcp(172.16.129.200:3307)/urm_resource?charset=utf8
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True", username, password, host, port, database, charset)
	db, err := gorm.Open(driverName, dsn)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	// 自动迁移模式 自动创建表
	db.AutoMigrate(&User{})
	return db
}

func isTelePhoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}
