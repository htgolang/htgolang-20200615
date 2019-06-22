package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createRandNum() int {
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(100)
	//fmt.Printf("rand is %v\n", rnd)
	return rnd
}

func main() {
	var input_num, rnd int
	var text string
	const max_num = 5
	rnd = createRandNum()

	for {

		for i := 1; i <= max_num; i++ {
			fmt.Printf("请输入一个0-100的整数： ")
			fmt.Scan(&input_num)

			if input_num > 100 || input_num < 0 {
				fmt.Printf("请输入0-100内的整数。")
			}

			if input_num > rnd {
				fmt.Printf("太大了，还有 %d 次机会。\n", max_num-i)
			} else if input_num < rnd {
				fmt.Printf("太小了，还有 %d 次机会。\n", max_num-i)
			} else if input_num == rnd {
				fmt.Printf("恭喜你猜对了！！！\n")
				break

			}

			if i == max_num {
				fmt.Printf("数字是： %d\n", rnd)
				fmt.Printf("真笨！！！\n")
			}

		}


		fmt.Printf("是否重新开始(y|n)  :")
		fmt.Scan(&text)
		// 判断用户输入
		if text == "n" {
			fmt.Printf("退出!")
			break
		}
	}
}

/*
 评分: 7
*/
