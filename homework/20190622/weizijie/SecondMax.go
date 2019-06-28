package main

import "fmt"

func main() {
	// 定义一个整数切片
	int_slice := []int{1, 4, 3, 1, 9, 6, 14}

	// 定义整数变量
	var max, second int
	// 将切片的前两个元素进行比较并赋值给max和second
	if int_slice[0] < int_slice[1] {
		max, second = int_slice[1], int_slice[0]
	}

	//从切片第二个元素与max和second分别进行比较，并赋值
	if len(int_slice) > 2 {
		for i := 2; i < len(int_slice); i++ {
			if int_slice[i] > second && int_slice[i] < max {
				second = int_slice[i]
			} else if int_slice[i] > max {
				max, second = int_slice[i], max
			}
		}
	}
	fmt.Printf("第二大的元素是: %d", second)

	/*
		// for遍历切片，循环比较两个数大小
		for i := 0; i < len(int_slice)-1; i++ {
			for j := 0; j < len(int_slice)-1-j; j++ {
				if int_slice[j] > int_slice[j+1] {
					int_slice[j], int_slice[j+1] = int_slice[j+1], int_slice[j]
				}

			}
		}
		fmt.Printf("int_slice 切片中最大的数：%d", int_slice[len(int_slice)-2])
	*/
}

/*
 评分: 8
*/
