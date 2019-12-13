package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main(){
	//启动一个例程
	go func(){
		for now := range time.Tick(1 * time.Second){
			fmt.Println(now)
		}
	}()

	//创建一个管道，长度为1
	ch := make(chan os.Signal, 1)
	//往管道里放入信号
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("wait")
	//当管道读取到数据时
	<-ch
	fmt.Println("over")
}