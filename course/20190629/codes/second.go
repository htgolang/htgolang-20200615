package main

import "fmt"

func main() {
	// 冒泡 循环两次 最后第二个索引的元素
	// 猴子掰苞谷 记录最大元素 每次比较元素与最大值关系 比最大值大交换
	// 第一个最大值和第二个最大值  初始化 [0]
	// 30 30
	nums := []int{1, 3, 5, 20, 30, 8, 9, 30, 11}
	maxNum, secondNum := nums[0], nums[0]

	for i, v := range nums {
		if v > maxNum {
			secondNum = maxNum
			maxNum = v
		} else if v == maxNum {
			continue
		} else if v > secondNum {
			secondNum = v
		}
		fmt.Println(i, ":", maxNum, secondNum)
	}
	fmt.Println(maxNum, secondNum)
}
