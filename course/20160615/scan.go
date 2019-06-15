package main

import "fmt"

func main() {
	var name string
	fmt.Print("请输入名字:")
	fmt.Scan(&name)
	fmt.Println("名字:", name)

	var age int
	fmt.Print("请输入你的年龄: ")
	fmt.Scan(&age)
	fmt.Println("年龄:", age)

	var height float64

	fmt.Print("请输入你的身高: ")
	fmt.Scan(&height)
	fmt.Println("身高:", height)
}
