package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置随机数种子
	rand.Seed(time.Now().Unix())
	const maxGuessNum = 5
	var isOk bool
	var guess int
	var input int

	for {
		//生成随机数0 - 100
		// guess := rand.Int() % 100
		guess = rand.Intn(100)
		for i := 0; i < maxGuessNum; i++ {
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
		}
		if !isOk {
			fmt.Println("太笨了, 游戏结束")
		}
		var text string
		fmt.Print("重新开始游戏吗?(y/n)")
		fmt.Scan(&text)
		if text != "y" {
			break
		}
	}

}
