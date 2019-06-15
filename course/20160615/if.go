package main

import "fmt"

func main() {
	// yes == "Y" "y"
	var yes string
	fmt.Print("有卖西瓜的吗?(Y/N):")
	fmt.Scan(&yes)

	fmt.Println("老婆的想法:")
	fmt.Println("十个包子")

	if yes == "Y" || yes == "y" {
		fmt.Println("一个西瓜")
	}

	fmt.Println("老公的想法:")

	if yes == "Y" || yes == "y" {
		fmt.Println("一个包子")
	} else {
		fmt.Println("十个包子")
	}

	var score int
	fmt.Print("请输入成绩:")
	fmt.Scan(&score)

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}
