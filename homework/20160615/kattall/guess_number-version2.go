package main

/*
	猜数字游戏
	1. 系统随机生成1-100的随机数
	2. 用户猜5次
		猜大了, 提示大了, 提示还有几次机会.
		猜小了, 提示小了, 提示还有几次机会.
		猜对了, 提示对了, 结束游戏. 游戏继续开始。
		全没猜对, 提示太弱了. 游戏继续开始。
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func randNumVerson2(num int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(num)
}

func gameStartVerson2() {
	// 循环让用户玩游戏
	for {
		// 生成随机数
		random_number := randNumVerson2(100)
		// fmt.Println("random_number: ", random_number)

		// 只能猜五次
		max := 5
		count := 1
		fmt.Println("随机数已经生成, 猜数字游戏开始.")
		for {
			// 让用用户输入
			fmt.Print("请输入1-100以内的整数: ")
			var num int
			fmt.Scan(&num)

			// 判断用户输入是否等于随机数
			if num > random_number {
				fmt.Printf("您猜大了. 您还有%d次机会.\n", max-count)
			} else if num < random_number {
				fmt.Printf("您猜小了. 您还有%d次机会.\n", max-count)
			} else {
				fmt.Println("恭喜您, 猜对了. 你太棒了.")
				break
			}

			if count == max {
				fmt.Println("您的次数已经用完, 游戏结束。")
				fmt.Println()
				break
			}

			count += 1
		}
	}
}
func main() {
	gameStartVerson2()
}
