package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 6, 5, 4}
	len_nums := len(nums)
	for i := 1; i < len_nums; i++ {
		tmp := nums[i]
		for j := i - 1; j >= 0; j-- {
			fmt.Println(tmp, nums[j])
			if tmp < nums[j] {
				nums[j+1], nums[j] = nums[j], tmp
			} else {
				break
			}
		}
	}
	fmt.Println(nums)
}

/*
	评分: 8
*/
