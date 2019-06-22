package main

import "fmt"

func main() {
	const NAME string = "kk"

	// 省略类型
	const PI = 3.1415926
	// 定义多个常量(类型相同)
	const C1, C2 int = 1, 2
	// 定义多个常量(类型不相同)
	const (
		C3 string = "silence"
		C4 int    = 1
	)
	// 定义多个常量 省略类型
	const C5, C6 = "silence", 1

	fmt.Println(NAME)
	fmt.Println(PI)
	fmt.Println(C1, C2)
	fmt.Println(C3, C4)
	fmt.Println(C5, C6)

	const (
		C7 int = 1
		C8
		C9
		C10 float64 = 3.14
		C11
		C12
		C13 string = "kk"
	)
	fmt.Println(C7, C8, C9)
	fmt.Println(C11, C12, C13)

	// 枚举, const+iota
	const (
		E1 int = iota
		E2
		E3
	)

	const (
		E4 int = iota
		E5
		E6
	)
	fmt.Println(E1, E2, E3)
	fmt.Println(E4, E5, E6)
}
