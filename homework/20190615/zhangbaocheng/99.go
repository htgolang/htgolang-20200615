package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d ", j, i, (j * i))
		}
		fmt.Println("")
	}
}

/*
 评分: 8
 尝试: 使用fmt.Println(""), fmt.Println()
	   注意代码格式
*/
