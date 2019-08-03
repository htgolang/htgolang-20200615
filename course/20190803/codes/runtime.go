package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.GOROOT())

	fmt.Println(runtime.NumCPU())

	fmt.Println(runtime.GOMAXPROCS(1))

	go func() {
		time.Sleep(3 * time.Second)
		runtime.Gosched()
	}()

	fmt.Println(runtime.NumGoroutine())
}
