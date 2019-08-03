package main

import (
	"fmt"
	"runtime"
	"time"
)

func PrintChars(name string) {
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%s: %c\n", name, ch)
		runtime.Gosched()
		// time.Sleep(time.Second)
	}
}

func main() {
	go PrintChars("1")
	go PrintChars("2")

	PrintChars("3")
	time.Sleep(time.Second * 3)
}
