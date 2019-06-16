package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	num := rand.Int() % 101
	// fmt.Println(num)

	for j := 5; j >= 1; j-- {
		var answer int
		fmt.Print("请输入一个0到100的整数: ")
		fmt.Scan(&answer)
		if answer == num {
			fmt.Println("恭喜,猜对了!")
			goto END
		} else if answer > num {
			if j == 1 {
				fmt.Println("没有猜对,真笨!")
				goto END
			} else {
				fmt.Printf("输入的数字太大了,请重新输入,还剩%d次机会\n", j-1)
			}
		} else if answer < num {
			if j == 1 {
				fmt.Println("没有猜对,真笨!")
				goto END
			} else {
				fmt.Printf("输入的数字太小了,请重新输入,还剩%d次机会\n", j-1)
			}
		}
	}
END:
}

