package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func addN(a, b int, args ...int) int {
	return 0
}

// callback 格式化 将传递的数据按照每行打印还是按照一行按|分割打印
func print(callback func(...string), args ...string) {
	fmt.Println("print函数输出:")
	callback(args...)
}

func list(args ...string) {
	for i, v := range args {
		fmt.Println(i, ":", v)
	}
}

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", addN)

	var f func(int, int) int = add

	fmt.Printf("%T\n", f)
	fmt.Println(f(1, 4))

	print(list, "A", "C", "E")

	//匿名函数
	sayHello := func(name string) {
		fmt.Println("Hello ", name)
	}

	sayHello("KK")
	sayHello("萌新")
	sayHello("武鹏飞")

	func(name string) {
		fmt.Println("Hi", name)
	}("KK")

	values := func(args ...string) {
		for _, v := range args {
			fmt.Println(v)
		}
	}
	a := "A"

	print(values, a, "B", "C")

	print(func(args ...string) {
		for _, v := range args {
			fmt.Print(v, "\t")
		}
		fmt.Println()
	}, "A", "C", "E")
}
