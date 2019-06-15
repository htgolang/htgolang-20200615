package main

import "fmt"

func main() {
	// i控制行，j控制列
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %2d  ", j, i, j*i)
		}
		// 每行打印完毕后，打印一个换行
		fmt.Println()
	}
}
