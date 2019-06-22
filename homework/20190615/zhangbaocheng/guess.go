package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

START:
	rand.Seed(time.Now().Unix())
	num := rand.Int() % 101
	fmt.Println(num)
	for j := 5; j >= 1; j-- {
		var input_num int
		fmt.Print("请输入一个0-100的整数:")
		fmt.Scan(&input_num)
		if input_num == num {
			fmt.Println("恭喜你，你猜对了！你是我肚子里的蛔虫吗？")
			goto END
		} else if input_num > num {
			if j == 1 {
				fmt.Println("你太笨了，这都没猜到, 游戏重新开始 ！")
				goto START
			} else {
				fmt.Println("大啦，大啦，大啦")
			}

		} else if input_num < num {
			if j == 1 {
				fmt.Println("你太笨了，这都没猜到,游戏重新开始 ！")
				goto START
			} else {
				fmt.Println("小啦，小啦，小啦")

			}

		}
	}

END:
}

/*
	评分: 7
	优化: 考虑在if else中相同逻辑是如何合并
*/
