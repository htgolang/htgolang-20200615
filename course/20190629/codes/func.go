package main

import "fmt"

// 定义(无参，无返回值)
func sayHelloWorld() {
	fmt.Println("Hello World!!!")
}

// 定义有参函数（形参）
func sayHi(name string) {
	fmt.Println("你好:", name)
}

func add(a int, b int) int {
	return a + b
}

func main() {
	// 打印标识符 sayHelloWorld类型
	fmt.Printf("%T\n", sayHelloWorld)

	// 调用函数
	sayHelloWorld()

	sayHi("kk") //实参
	name := "祥哥"
	sayHi(name) //实参name值(祥哥) 传递给 函数的形参name

	rt := add(1, 5)
	fmt.Println(rt)

}
