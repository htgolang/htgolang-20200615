package main

import "fmt"

func main() {
	number := []int{10, 20, 30, 40, 55, 66, 77, 88, 99, 100}
	// 定义切片中第一个元素和最后一个元素的位置及中间元素的位置
	min := 0
	max := len(number) - 1
	var middle int

	// 从键盘输入一个整数
	var num int
	fmt.Println("请输入100以内的整数: ")
	fmt.Scan(&num)

	// for循环，依次匹配切片中的元素，找到匹配的值或相近的的值则退出
	for {
		// 定义切片的中间索引位置
		middle = (min + max) / 2

		// 如果输入的值大于切片的中间值，则判断该值是否在中间值和中间后第一个值之间，若是，则返回相近的值，若不是，则下一次循环的起始索引为middle+1(中间索引的后一个索引)
		if num > number[middle] {
			if num < number[middle+1] {
				fmt.Printf("没有匹配的数，但你的输入是在%d - %d之间", number[middle], number[middle+1])
				break
			} else {
				min = middle + 1
			}
			// 如果输入的值小于切片的中间值，则判断该值是否在中间值和中间值的前一个索引值之间，若是，则返回相近的值，若不是，则下一次循环的结束索引为middle-1(中间索引的前一个索引)
		} else if num < number[middle] {
			if num > number[middle-1] {
				fmt.Printf("没有匹配的数，但你的输入是在%d - %d之间", number[middle-1], number[middle])
				break
			} else {
				max = middle - 1
			}
			// 如果该值匹配到切片中某一个元素，则返回该元素的索引值
		} else {
			fmt.Printf("对应匹配的位置是在第%d个元素，匹配的值为:%v", middle, number[middle])
			break
		}
	}
}

/*
 评分: 6
 二分查找算法 用于对数据的查询
*/
