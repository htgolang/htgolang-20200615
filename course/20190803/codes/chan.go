package main

import "fmt"

func main() {
	var notice chan string = make(chan string)

	fmt.Printf("%T, %v\n", notice, notice)

	go func() {
		fmt.Println("go start")
		notice <- "xxxx"
		fmt.Println("go end")
	}()

	fmt.Println("start")
	fmt.Println(<-notice)
	fmt.Println("end")
}
