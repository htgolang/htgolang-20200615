package main

import "fmt"

// 传入一个切片, 返回一个最大的整数
func findMaxInt(numbers []int) int {

	// 冒泡排序, 只排一次, 取最大值即可。
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
		}
	}
	return numbers[len(numbers)-1]
}

func main() {
	/*
		找int切片中最大的元素(不准用排序)
	*/

	// 定义一个slince int类型
	numbers := []int{6, 8, 3, 1, 4, 9, 2}

	// 调用函数findMaxInt   max接受返回值
	max := findMaxInt(numbers)

	// 输出
	fmt.Println("最大的数是：", max)

}
