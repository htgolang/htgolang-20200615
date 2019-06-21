package main

import "fmt"

func main() {
	fmt.Println("打印乘法口诀表：")
	for d1 := 1; d1 <= 9; d1++ {
		for d2 := 1; d2 <= 9; d2++ {
			if d2 > d1 {
				break
			}
			fmt.Printf("%d X %d = %2d  ", d2, d1, d1*d2)
		}
		fmt.Println()
	}
}

/*
 评分: 7
 思考: 考虑不使用if break判断如何实现
*/
