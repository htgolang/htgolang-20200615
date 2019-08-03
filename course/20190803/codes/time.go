package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now())

	fmt.Println("after")
	channel := time.After(3 * time.Second)
	fmt.Println(<-channel) // 只读一次 //延迟执行一次

	fmt.Println("ticker")
	ticker := time.Tick(3 * time.Second)
	fmt.Println(<-ticker) //每隔ns产生一个管道消息 //每隔ns执行动作
	fmt.Println(<-ticker)
	fmt.Println(<-ticker)

	for now := range ticker {
		fmt.Println(now)
	}
}
