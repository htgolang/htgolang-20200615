package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var count int = 5
	var gameRes bool

	for {
		var gameSwitch string
		fmt.Printf("猜数游戏即将开始(共有5次机会), 请确认是否继续<yes/no>")
		fmt.Scan(&gameSwitch)

		//控制玩家只能输入yes或no
		switch gameSwitch {
		case "yes":
			break
		case "no":
			break
		default:
			fmt.Println("请输入yes or no")
			continue
		}

		//如果是no, 结束游戏
		if gameSwitch == "no" {
			break
		}

		//生成0-100随机数
		rand.Seed(time.Now().Unix())
		randNum := rand.Intn(100 - 0)
		randNum = randNum + 0
		fmt.Printf("谜底是: %d\n", randNum)

		//game主程序
		for i := 1; i <= 5; i++ {
			var inputNum int
			fmt.Printf("请输入猜数: ")
			fmt.Scan(&inputNum)
			residue := count - i
			if inputNum > randNum {
				fmt.Printf("猜数太大了, 你还有%d次机会\n", residue)
			} else if inputNum < randNum {
				fmt.Printf("猜数太小了, 你还有%d次机会\n", residue)
			} else {
				fmt.Println("恭喜你答对了")
				gameRes = true
				break
			}
			if i == 5 {
				gameRes = false
			}
		}

		//控制是否重新开始游戏
		if gameRes == false {
			fmt.Println("菜鸡, 让我们重新开始游戏吧.")
		} else {
			break
		}
	}
}

/*
	评分: 8
*/
