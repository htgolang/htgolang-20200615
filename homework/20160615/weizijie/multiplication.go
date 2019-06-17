package main

import "fmt"

// 打印九九乘法表
func main() {
	for j := 1; j <= 9; j++ {
		for i := 1; i <= j; i++ {
			fmt.Printf("%d * %d = %-2d  ", i, j, i*j)
		}
		fmt.Println("")
	}
}
