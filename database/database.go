package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"myblog/model"
)

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
	db.AutoMigrate(&model.User{})
	return db
}
