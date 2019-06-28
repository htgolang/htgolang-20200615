package main

import "fmt"

// 传入一个切片, 返回一个最大的整数
func findMaxSecondInt(numbers []int, maxSeond int) int {

	// 冒泡排序
	for j := 0; j < len(numbers); j++ {
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
			}
		}
	}
	return numbers[len(numbers)-maxSeond]
}

func main() {
	/*
		找int切片中第二个最大的元素(不准用排序)
	*/

	// 定义一个slince int类型
	numbers := []int{6, 8, 3, 1, 4, 9, 2}

	// 第二个最大的元素
	maxSeond := 2

	// 调用函数findMaxInt   max接受返回值
	max := findMaxSecondInt(numbers, maxSeond)

	// 输出
	fmt.Println("第二个最大的元素是：", max)

}

/*
 评分: 8
*/
