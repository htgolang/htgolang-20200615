package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("new")
			return 1
		},
	}

	x := pool.Get()

	fmt.Println(x)
	pool.Put(x)

	x = pool.Get()
	x = pool.Get()
}
