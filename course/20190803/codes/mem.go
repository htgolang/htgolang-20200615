package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int
	group := &sync.WaitGroup{}
	lock := &sync.Mutex{}

	incr := func() {
		defer group.Done()
		for i := 0; i <= 100; i++ {
			lock.Lock()   // 加锁（互斥锁）
			counter++     //a1.拿counter(0)  a2.counter+1 a3.存counter 1
			lock.Unlock() //释放错

			runtime.Gosched()
		}
	}

	decr := func() {
		defer group.Done()
		for i := 0; i <= 100; i++ {
			lock.Lock()
			counter-- //b1.拿counter(0)  b2.counter-1 b3.存counter
			lock.Unlock()
			runtime.Gosched()
		}
	}

	//a1 b1 a2 a3(1) b2 b3(-1)

	for i := 0; i < 10; i++ {
		group.Add(2)
		go incr()
		go decr()
	}

	group.Wait()
	fmt.Println(counter)
}
