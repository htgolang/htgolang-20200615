package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//提前声明变量，防止循环声明引起异常
	var i int
	var ReadNum int

	for {
		//逻辑锁在死循环里
		//程序起点，简称打开游戏，以下都为反复循环角度考虑
	STARTEND:
		//重置随机数
		rand.Seed(time.Now().Unix())
		RandNum := rand.Int() % 100
		//打印出随机数，方便测试数字大小的提示。
		//fmt.Println(RandNum)

		//开始猜数主体
		for i = 5; i >= 0; i-- {
			//调试i的值
			//fmt.Println(i)
			//判断游戏何时退出
			if i != 0 {
				//开始正常游戏过程
				fmt.Printf("请猜数字，你还有%d次机会，加油！！: ", i)
				fmt.Scan(&ReadNum)
				switch {
				case ReadNum > RandNum:
					fmt.Println("有点大了")
				case ReadNum < RandNum:
					fmt.Println("有点小了")
				case ReadNum == RandNum:
					fmt.Println("恭喜你，猜中了！！")
					fmt.Println("请重新开始猜数游戏。。")
					//重新打开游戏
					goto STARTEND
				}
			} else {
				fmt.Println("你太笨了")
				//重新打开游戏
				goto STARTEND

			}

		}

	}
}

/*
 评分: 8
 改进: 随机种子在程序中只用设置一次
*/
