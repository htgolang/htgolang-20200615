package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/*
		2. 猜数字游戏
			a. 生成一个[0,100)随机数
			b. 让用户最多猜5次(从命令行输入一个整数)
				猜的太大 => 提示太大了，你还有N次猜测机会
				猜的太小 => 提示太小了，你还有N次猜测机会
				猜中了 => 猜中了

				5次都没猜中 => 退出，并提示太笨了
		挑战:
			5次都没有猜中，提示太笨了，重新开始猜数游戏(重新生成随机数，重新开始5次计数)
	*/

	var randnum, usernum int
	randnum = rand.Intn(100)
	fmt.Println("随机数为: ", randnum)
	// fmt.Println(randnum)
	for i := 1; i <= 5; i++ {
		fmt.Print("请输入一个整数:")
		fmt.Scan(&usernum)
		if randnum == usernum {
			fmt.Println("猜中了")
			goto END
		} else if randnum < usernum {
			fmt.Println("猜的太大")
		} else {
			fmt.Println("猜的太小")
		}
	}
	fmt.Println("太笨了！！！")
END:
}
