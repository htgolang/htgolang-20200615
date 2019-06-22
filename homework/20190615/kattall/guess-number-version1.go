package main

/*
	猜数字游戏
	1. 系统随机生成1-100的随机数
	2. 用户猜5次
		猜大了, 提示大了, 提示还有几次机会.
		猜小了, 提示小了, 提示还有几次机会.
		猜对了, 提示对了, 结束游戏.
		全没猜对, 提示太弱了
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func randNumVerson1(num int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(num)
}

func gameStartVerson1(max int) {
	// 生成随机数
	random_number := randNumVerson1(100)
	// fmt.Println("random_number: ", random_number)

	// 猜的次数
	count := 1

	fmt.Println("猜数字游戏开始.")
	for {
		// 让用用户输入
		fmt.Print("请输入1-100以内的整数: ")
		var num int
		fmt.Scan(&num)

		// 判断用户输入是否等于随机数
		if num > random_number {
			// 判断是否已经结束, 已经结束直接结束游戏.
			if max == count {
				fmt.Println("您的次数已经用完,你太弱了,游戏结束。")
				goto END
			}
			fmt.Printf("您猜大了. 您还有%d次机会.\n", max-count)

		} else if num < random_number {
			// 判断是否已经结束, 已经结束直接结束游戏.
			if max == count {
				fmt.Println("您的次数已经用完,你太弱了,游戏结束。")
				goto END
			}
			fmt.Printf("您猜小了. 您还有%d次机会.\n", max-count)
		} else {
			fmt.Println("恭喜您, 猜对了. 你太棒了.")
			return
		}
		count += 1
	}
END:
}

func main() {
	// 传入参数, 最多失败的次数。
	gameStartVerson1(5)
}
