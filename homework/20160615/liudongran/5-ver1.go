package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := rand.Intn(100)
	//fmt.Printf(a)
	for c := 0; c <=5; c++ {
		fmt.Printf("请输入1-100之间的数 \n")
		var num int
		fmt.Scan(&num)
		if num > a {
			fmt.Printf("你猜大了，你还有%d次猜测机会 \n", 5-c)
		} else if num < a {
			fmt.Printf("你猜小了，你还有%d次猜测机会 \n", 5-c)
		} else {
			fmt.Printf("你猜对了 \n")
			break
		}
	if c == 5{
			fmt.Print("你好笨啊！\n")
		}
	}
}
