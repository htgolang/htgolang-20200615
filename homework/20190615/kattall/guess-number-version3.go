package main

/*
	猜数字游戏
	1. 系统随机生成1-100的随机数
	2. 用户猜5次
		猜大了, 提示大了, 提示还有几次机会.
		猜小了, 提示小了, 提示还有几次机会.
		猜对了, 提示对了, 结束游戏. 提示是否继续游戏, 输入y继续游戏。输入其他任意字符，结束游戏。
		全没猜对, 提示太弱了. 提示是否继续游戏, 输入y继续游戏。输入其他任意字符，结束游戏。
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func randNumVerson3(num int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(num)
}

func gameStartVerson3(max int) {
	// 循环让用户玩游戏
	for {
		// 生成随机数
		random_number := randNumVerson3(100)
		// fmt.Println("random_number: ", random_number)

		// 猜的次数
		count := 1
		fmt.Println("随机数已经生成, 猜数字游戏开始.")
		for {
			// 让用用户输入
			fmt.Print("请输入1-100以内的整数: ")
			var num int
			fmt.Scan(&num)

			// 判断用户输入是否等于随机数
			if num > random_number {
				// 判断是否已经结束, 已经结束直接结束游戏.
				if max == count {
					fmt.Println("您的次数已经用完,你太弱了,游戏结束。\n")
					goto IS_AGAIN
				}
				fmt.Printf("您猜大了. 您还有%d次机会.\n", max-count)
			} else if num < random_number {
				// 判断是否已经结束, 已经结束直接结束游戏.
				if max == count {
					fmt.Println("您的次数已经用完,你太弱了,游戏结束。\n")
					goto IS_AGAIN
				}
				fmt.Printf("您猜小了. 您还有%d次机会.\n", max-count)
			} else {
				fmt.Println("恭喜您, 猜对了. 你太棒了.")
				break
			}

		IS_AGAIN:
			if count == max {
				fmt.Print("请问你还要继续玩吗? (Y/N)")
				var again string
				fmt.Scan(&again)
				if again == "Y" || again == "y" {
					break
				} else {
					return
				}
			}
			count += 1
		}
	}
}

func main() {
	// 传入参数, 最多失败的次数。
	gameStartVerson3(5)
}

/*
	评分: 6
	优化: 考虑在if else中相同逻辑是如何合并
*/
