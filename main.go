package main

import (
	"myblog/router"
	"os"
	"os/signal"
)

func main() {

	go router.StartProxy()

	//监听主程序是否结束
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	select {
	case <-c:
		break
	}

}
