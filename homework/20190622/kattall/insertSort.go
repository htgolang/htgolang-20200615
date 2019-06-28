package main

import "fmt"

func sortInt(values []int) []int {
	n := len(values)

	// 如果切片长度是1 直接返回
	if n < 2 {
		return values
	}

	// 插入排序
	for i := 1; i < n; i++ {
		for j := i; j > 0 && values[j] < values[j-1]; j-- {
			values[j], values[j-1] = values[j-1], values[j]
		}
	}

	return values
}

func main() {
	nums := []int{3, 7, 5, 9, 10, 4}
	fmt.Println("插入排序后的切片为：", nums)

	sortNums := sortInt(nums)
	fmt.Println("插入排序后的切片为：", sortNums)
}

/*
 评分: 8
*/
