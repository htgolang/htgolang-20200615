package main

import (
	"fmt"
	"runtime"
)

func PrintChars(name int, channel chan int) {
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%d: %c\n", name, ch)
		runtime.Gosched()
	}
	channel <- name
}

func main() {
	var channel chan int = make(chan int)

	for i := 0; i < 10; i++ {
		go PrintChars(i, channel)
	}

	for i := 0; i < 10; i++ {
		<-channel
	}

	fmt.Println("over")
}
