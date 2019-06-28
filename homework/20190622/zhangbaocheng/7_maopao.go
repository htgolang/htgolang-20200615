package main

import "fmt"

func main() {
	nums := []int{6, 5, 4, 1, 2, 3}
	len_num := len(nums)
	for i := 0; i < len_num; i++ {
		for j := 0; j < i; j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
			fmt.Println(nums)
		}
	}
	fmt.Println(nums)

}

/*
	评分: 8
*/
