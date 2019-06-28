package main

import "fmt"

func main() {
	nums := []int{6, 5, 4, 1, 2, 3}
	max_num := 0
	for _, num := range nums {
		if num > max_num {
			max_num = num
		}
	}
	fmt.Printf("最大数为: %d", max_num)

}

/*
	评分: 8
*/
