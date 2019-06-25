package main

import "fmt"

func main() {
	// 定义一个整数切片
	int_slice := []int{1, 4, 2, 7, 5, 3}

	// for遍历切片，循环比较两个数大小
	for i := 0; i < len(int_slice)-1; i++ {
		if int_slice[i] > int_slice[i+1] {
			int_slice[i], int_slice[i+1] = int_slice[i+1], int_slice[i]
		}
	}
	fmt.Printf("int_slice 切片中最大的数：%d", int_slice[len(int_slice)-1])

}
