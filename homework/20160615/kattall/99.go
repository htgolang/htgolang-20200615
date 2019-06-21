package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d \t", i, j, i*j)
		}
		fmt.Println(" ")
	}
}

/*
 评分: 8
 尝试: 使用fmt.Println(""), fmt.Println()
*/
