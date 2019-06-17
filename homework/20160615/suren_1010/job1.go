package main

import "fmt"

func main() {

	/* 简洁版
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}*/

	// 对齐版
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			if i > 2 && i <= 4 && j == 2 {
				fmt.Printf("%dx%d=%d  ", j, i, i*j)
			} else {
				fmt.Printf("%dx%d=%d ", j, i, i*j)
			}
		}
		fmt.Println()
	}
}
