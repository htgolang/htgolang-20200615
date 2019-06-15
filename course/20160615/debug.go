package main

import "fmt"

func main() {
	var age = 30
	age = age + 1
	age = age + 1
	fmt.Println("后年：", age)

	fmt.Println("打印第一行")
	fmt.Println("打印第2行")
	fmt.Print("打印第一行")
	fmt.Print("打印第2行")
	fmt.Printf("\n%T, %s, %T, %d\n", "KK", "KK", age, age)
}
