package main

import "fmt"

func main() {

	addBase := func(base int) func(int) int {
		// 返回函数
		return func(n int) int {
			return base + n
		}
	}

	add10 := addBase(10)
	add100 := addBase(100)

	fmt.Println(addBase(2)(3))
	fmt.Println(add10(3))
	fmt.Println(add100(3))

	add8 := addBase(8)

	fmt.Printf("%T\n", add8)
	fmt.Println(add8(10))
	fmt.Println(add8(20))

}
