package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
START: // 重新开始的点
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(100)

	var inputNum int
	// debug, 打印随机生成的那个数，用来验证代码是否正确
	fmt.Printf("This number is %d\n", randNum)

	// 循环五次
	for i := 1; i <= 5; i++ {
		fmt.Print("The range of number in 0-100。Guess which it is?:")
		fmt.Scan(&inputNum)

		// 判断大小
		switch {
		case inputNum == randNum:
			// 输出猜中了，并跳转至END结束
			fmt.Printf("Bingo! The number is %d\n", randNum)
			goto END
		case inputNum > randNum:
			// 5次机会用光，则跳转至START，重新开始。否则输出信息后，继续进行下一次猜数
			if i == 5 {
				fmt.Print("You're too stupid, Game Restart\n\n\n")
				goto START
			} else {
				fmt.Print("Input number it's too big, ")
			}

		case inputNum < randNum:
			// 5次机会用光，则跳转至START，重新开始。否则输出信息后，继续进行下一次猜数
			if i == 5 {
				fmt.Print("You're too stupid, Game Restart\n\n\n")
				goto START
			} else {
				fmt.Print("Input number it's too small, ")
			}

		}

		// 每一次猜数结束后，利用5-i 计算剩余次数
		if 5-i != 1 {
			fmt.Printf("You still have %d chances\n\n", 5-i)
		} else {
			fmt.Print("You still have 1 chance\n\n")
		}

	}
END:
}
