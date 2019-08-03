package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int32 = 0

	group := &sync.WaitGroup{}

	incr := func() {
		defer group.Done()
		for i := 0; i <= 100; i++ {
			atomic.AddInt32(&counter, 1)
			runtime.Gosched()
		}
	}

	decr := func() {
		defer group.Done()
		for i := 0; i <= 100; i++ {
			atomic.AddInt32(&counter, -1)
		}
	}

	for i := 0; i < 10; i++ {
		group.Add(2)
		go incr()
		go decr()
	}

	group.Wait()
	fmt.Println(counter)
}
