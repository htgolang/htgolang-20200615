package main

import "fmt"

func main() {
	// 冒泡 循环一次 最后一个索引的元素
	// 猴子掰苞谷 记录最大元素 每次比较元素与最大值关系 比最大值大交换
	// 初始化 [0] 切片/数组 元素数量 0

	nums := []int{1, 3, 5, 20, 30, 8, 9, 11}
	maxNum := nums[0]

	for i, v := range nums {
		if v > maxNum {
			maxNum = v
		}
		fmt.Println(i, ":", maxNum)
	}
	fmt.Println(maxNum)
}
