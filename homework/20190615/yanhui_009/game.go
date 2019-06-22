/*
题目二:
2、猜数字游戏
	a. 生成一个随机数(0~100的随机整数)
	b. 让用户重复猜5次，(用户从命令行键入猜的整数)
		当猜的太大 => 提示太大了，并提示你还有几次猜测机会
		当猜的太小 => 提示太小了，并提示你还有几次猜测机会
		猜中了 => 提示猜中了

		如果5次都没有猜中 => 退出，并提示太笨了
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // 产生随机生成随机数的种子
	var num int = rand.Intn(100) // 生成一个100以内的随机正整数
	//fmt.Println(num)
	var guess_num int     // 定义用户猜测数据的变量
	const total_count = 5 //定义猜游戏的总次数

	//fmt.Scan(&guess_num)
	for count := 1; count <= 5; count++ {
		fmt.Printf("请输入100以内任意一个正整数:")
		fmt.Scan(&guess_num)
		if guess_num > num {
			if count == total_count {
				//fmt.Printf("您猜测的数字太大了，您本轮猜字游戏次数用尽。\n")
				fmt.Println("您太笨啦! 我们本轮猜数游戏的数字为:", num)
				break
			}
			fmt.Printf("抱歉,您本次猜测的数字太大。本轮猜字游戏您还剩余%d次机会\n", total_count-count)
			continue
		} else if guess_num < num {
			if count == total_count {
				//fmt.Printf("您猜测的数字太小了，您本轮猜字游戏次数用尽。\n")
				fmt.Println("您太笨啦! 我们本轮猜数游戏的数字为:", num)
				break
			}
			fmt.Printf("抱歉,您本次猜测的数字太小。本轮猜字游戏您还剩余%d次机会\n", total_count-count)
			continue
		} else {
			fmt.Println("恭喜你，猜对了。猜字游戏生成的答案为:", num)
			break
		}
	}

}

/*
 评分: 6
 思考： continue用途
        if else中相同的逻辑如何合并
*/
