package main

import (
	"fmt"
	"sort"
)

func main() {
	//  排序
	nums := []int{1, 3, 5, 9, 8, 7, 4}
	sort.Ints(nums)
	fmt.Println(nums)

	// 二分查找 在有序的数组中查找元素

	fmt.Println(sort.SearchInts(nums, 5))
	fmt.Println(sort.SearchInts(nums, 6))

	num := 5
	if nums[sort.SearchInts(nums, num)] == num {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
}
