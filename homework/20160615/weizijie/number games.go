package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
START:

	// 随机获取100内整数
	rand.Seed(time.Now().Unix())
	var rangenumber int = rand.Intn(100)
	//fmt.Println("Jevon ",rangenumber)

	// 用户循环输入五次
	for i := 1; i <= 5; i++ {
		fmt.Println("---------------------- ")
		var number int
		fmt.Println("请输入一个100以内的整数: ")
		fmt.Scan(&number)
		//fmt.Printf("输入的number是: %d\n", number)

		// 判断用户输入的是否正确
		switch {
		case number == rangenumber:
			fmt.Printf("你输入的number为：%d,恭喜你，猜对了.", number)
			goto BREAKEND
		case number > rangenumber:
			fmt.Printf("你输入的number为：%d,很遗憾，你猜大了.\n你还剩%d次机会\n\n", number, 5-i)

		case number < rangenumber:
			fmt.Printf("你输入的number为：%d,很遗憾，你猜小了.\n你还剩%d次机会\n\n", number, 5-i)
		}

		// 如果用户输错了五次，则让用户选择是否重新开始
		if i == 5 {
			var yes string
			fmt.Println("你已经输错了5次，此轮已结束，你是否要直接开始下一轮,若直接开始下一轮，则输入Y或y,若直接退出，输入或其他: ")
			fmt.Scanln(&yes)
			fmt.Scan(&yes)

			if yes == "Y" || yes == "y" {
				fmt.Printf("输入的是: %s,则开始下一轮游戏\n\n", yes)
				goto START
			} else {
				fmt.Printf("输入的是: %s,则结束游戏", yes)
			}
		}
	}
BREAKEND:
}

/*
 评分: 8
*/
