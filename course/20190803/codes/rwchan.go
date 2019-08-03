package main

import (
	"fmt"
	"time"
)

func main() {
	var channel chan int = make(chan int, 5)

	var rchannel <-chan int = channel
	var wchannel chan<- int = channel

	go func(channel <-chan int) {
		fmt.Println(<-channel)
	}(channel)

	go func(channel chan<- int) {
		channel <- 0
	}(channel)

	wchannel <- 1
	fmt.Println(<-rchannel)
	time.Sleep(time.Second * 3)

}
