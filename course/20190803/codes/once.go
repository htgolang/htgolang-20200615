package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	for i := 0; i < 10; i++ {
		once.Do(func() {
			fmt.Println(i)
		})
	}
}
