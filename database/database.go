package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"myblog/model"
)

//数据库 gorm
func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.userName")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	//"mysql":      "mysql",    // root:Netposa123@tcp(172.16.129.200:3307)/urm_resource?charset=utf8
	//注意：dsn中要加loc=Local或者loc=Asia%2FShanghai，否则时区错误
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&%s", username, password, host, port, database, charset, loc)
	db, err := gorm.Open(driverName, dsn)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	// 自动迁移模式 自动创建表
	db.AutoMigrate(&model.User{})
	return db
}
