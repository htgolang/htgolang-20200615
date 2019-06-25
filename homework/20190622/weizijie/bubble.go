package main

import "fmt"

func main() {

	`冒泡排序`

	// 定义一个整型切片
	sort_int := []int{1, 6, 3, 7, 8, 36, 28, 4, 2}

	fmt.Printf("按从小到大排序之前的结果是: %v\n", sort_int)

	// 先比较前两个数的大小，依次向后移动比较
	for j := 0; j < len(sort_int)-1; j++ {
		for i := 0; i < len(sort_int)-1-j; i++ {
			if sort_int[i] > sort_int[i+1] {
				sort_int[i], sort_int[i+1] = sort_int[i+1], sort_int[i]
			}
		}
	}

	fmt.Printf("按从小到大排序之后的结果是: %v\n", sort_int)

}
