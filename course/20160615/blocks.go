package main

import "fmt"

func main() {
	// 作用域: 定义标识符可以使用的范围
	// 在Go中用{}来定义作用域的范围
	// 使用原则: 子语句块可以使用父语句块中的标识符, 父不能使用子的

	outer := 1
	{
		inner := 2
		fmt.Println(outer)
		fmt.Println(inner)
		outer := 21
		{
			inner2 := 3
			fmt.Println(outer, inner, inner2)
		}
	}
	fmt.Println(outer)
	// fmt.Println(inner)

}
