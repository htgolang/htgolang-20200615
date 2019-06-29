package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer01")
	}()
	defer func() {
		fmt.Println("defer02")
	}()

	fmt.Println("main over")
}
