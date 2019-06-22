package main

import "fmt"

func main() {
	// float32, float64
	// 字面量: 十进制表示法 科学技术表示法
	// M E N => M * 10 ^ N
	// 1.9 E -1 =>0.19
	var height float64 = 1.68
	fmt.Printf("%T %f\n", height, height)
	var weight float64 = 13.05E1
	fmt.Println(weight)

	// 操作
	// 算数运算(+, -, *, /, ++, --)
	fmt.Println(1.1 + 1.2)
	fmt.Println(1.1 - 1.2)
	fmt.Println(1.1 * 1.2)
	fmt.Println(1.1 / 1.2)

	height++
	fmt.Println(height)
	height--
	fmt.Println(height)

	// 关系运算(> >= < <=)
	fmt.Println(1.1 > 1.2)
	fmt.Println(1.1 >= 1.2)
	fmt.Println(1.1 < 1.2)
	fmt.Println(1.1 <= 1.2)

	fmt.Println(1.2-1.1 <= 0.005)
	// 赋值(=, +=, -=, /= *=)

	height += 0.05
	fmt.Println(height)

	fmt.Printf("%T, %T\n", 1.1, float32(1.1))

	fmt.Printf("%5.2f\n", height)
}
