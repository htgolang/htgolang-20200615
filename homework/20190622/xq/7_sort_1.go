package main

import "fmt"

// 冒泡排序

func main() {
	test_slice := []int{10, 50, -1, -3, 0}

	for i := 0; i <= len(test_slice)-1; i++ {
		for j := i + 1; j < len(test_slice); j++ {
			if test_slice[i] > test_slice[j] {
				test_slice[i], test_slice[j] = test_slice[j], test_slice[i]
			}
		}
	}

	fmt.Println(test_slice)
}

/*
	评分: 8
	思考代码执行流程及切片变化
*/
