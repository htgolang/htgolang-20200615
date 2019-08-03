package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	fmt.Println(len(channel))
	channel <- "x"
	fmt.Println(len(channel))
	channel <- "y"
	fmt.Println(len(channel))

	fmt.Println(<-channel)
	fmt.Println(len(channel))
	fmt.Println(<-channel)
	fmt.Println(len(channel))

	channel <- "z"
	channel <- "a"
	close(channel)

	// fmt.Println(<-channel)
	// channel <- "a"

	for ch := range channel { // 需要有某个例程能够关闭管道，否则会发生死锁
		fmt.Println(ch)
	}

}
