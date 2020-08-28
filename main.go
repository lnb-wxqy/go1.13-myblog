package main

import (
	"github.com/json-iterator/go/extra"
	"github.com/spf13/viper"
	"log"
	"myblog/router"
	"os"
	"os/signal"
)

func main() {

	// josn模糊匹配
	extra.RegisterFuzzyDecoders()

	//初始化配置
	InitConfig()

	go router.StartProxy()

	//监听主程序是否结束
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	select {
	case <-c:
		break
	}

}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("init application.yml failed, err: ", err)
		panic(err)
	}
}
