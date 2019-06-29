package main

import "fmt"

func main() {
	test_slice := []int{1, 0, 4, 3, -2, -1, 10}

	//var max_num int
	// 设定初始值为数组的第一个元素
	max_num := test_slice[0]

	for i := range test_slice {

		if test_slice[i] > max_num {
			max_num = test_slice[i]
		}
	}

	fmt.Printf("最大的值为：%d\n", max_num)
}

/*
	评分: 7
*/
