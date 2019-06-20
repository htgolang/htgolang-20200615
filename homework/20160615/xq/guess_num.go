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

func main(){
	var input_num, rnd int
	rnd = createRandNum()

	for i:=1;i<=5;i++ {
		fmt.Printf("请输入一个0-100的整数： ")
		fmt.Scan(&input_num)

		if input_num >100 || input_num < 0 {
			fmt.Printf("请输入0-100内的整数。")
			continue
		}

		if input_num > rnd {
			fmt.Printf("太大了，还有 %d 次机会。", 5-i)
		}else if input_num < rnd {
			fmt.Printf("太小了，还有 %d 次机会。", 5-i)
		}else if input_num == rnd {
			fmt.Printf("恭喜你猜对了！！！")
			break

		}

		if i == 5 {
			fmt.Printf("数字是： %d\n", rnd)
			fmt.Printf("真笨！！！")
		}

	}
}