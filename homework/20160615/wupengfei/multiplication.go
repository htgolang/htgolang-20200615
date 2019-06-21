package main

import "fmt"

func main() {
	// 1. 打印乘法口诀
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i == j {
				fmt.Printf("%dx%d=%d\n ", j, i, j*i)
				break
			}
			fmt.Printf("%dx%d=%d ", j, i, j*i)
		}
	}
}

/*
 评分: 7
 思考: 考虑不使用if break判断如何实现
*/
