package main

import "fmt"

func main()  {
	test_slice := []int{1, 0, 4, 3, -2, -1, 10}

	var max_num int

	for i := range test_slice {

		if test_slice[i] > max_num {
			max_num = test_slice[i]
		}
	}

	fmt.Printf("最大的值为：%d\n", max_num)
}
