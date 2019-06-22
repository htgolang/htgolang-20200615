package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置随机数种子
	rand.Seed(time.Now().Unix())

	// 生成随机数0 - 100
	// guess := rand.Int() % 100
	guess := rand.Intn(100)
	const maxGuessNum = 5
	var isOk bool
	for i := 0; i < maxGuessNum; i++ {
		var input int
		fmt.Print("请输入你猜的值:")
		fmt.Scan(&input)

		if guess == input {
			fmt.Printf("太聪明了, 你猜测%d次就猜对了\n", i+1)
			isOk = true
			break
		} else if input > guess {
			fmt.Printf("太大了, 你还有%d猜测机会\n", maxGuessNum-i-1)
		} else {
			fmt.Printf("太小了, 你还有%d猜测机会\n", maxGuessNum-i-1)
		}

		// // 当最后一次执行这里(最后一次也没有猜测成功)
		// if i == maxGuessNum-1 {
		// 	fmt.Println("太笨了, 游戏结束")
		// }
	}
	if !isOk {
		fmt.Println("太笨了, 游戏结束")
	}
}
