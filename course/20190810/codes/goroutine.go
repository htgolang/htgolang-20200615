package main

import (
	"fmt"
	"time"
)

func copySlice(dest []string, src []string) {
	for i := 0; i < len(src) && i < len(dest); i++ {
		dest[i] = src[i]
	}
}

func main() {
	names := []string{"kk", "xiaofan"}
	// dnames := make(map[string]string)

	// cnames := make([]string, 1)

	// copySlice(cnames, names)
	// fmt.Println(cnames, names)

	//copy(cnames, names)

	channel := make(chan []string)

	go func(channel chan []string) {
		names := <-channel
		names[0] = "0000"
	}(channel)

	// //names = append(names, "kk")
	// names[0] = "kk"
	// dnames["kk"] = "test"

	channel <- names
	time.Sleep(time.Second * 4)
	fmt.Println(names)
}
