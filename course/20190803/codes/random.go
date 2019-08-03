package main

import "fmt"

func main() {
	channel := make(chan int, 1)
	slice := make([]int, 10)

	for i := 0; i < 10; i++ {
		select {
		case channel <- 0:
		case channel <- 1:
		case channel <- 2:
		case channel <- 3:
		case channel <- 4:
		case channel <- 5:
		}

		// fmt.Println(<-channel)
		slice[i] = <-channel
		// slice = append(slice, <-channel)
	}
	fmt.Println(slice)
}
