package main

import "fmt"

func main() {

	// 插入排序
	test_slice := []int{10, 50, -1, -3, 0}

	for i := 0; i <= len(test_slice)-1; i++ {
		tmp := test_slice[i]

		for j := i - 1; j >= 0; j-- {
			if tmp < test_slice[j] {
				test_slice[j+1] = test_slice[j]
				test_slice[j] = tmp
			}
		}
	}
	fmt.Println(test_slice)

}

/*
	评分: 8
*/
