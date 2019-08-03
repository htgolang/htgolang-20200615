package main

import (
	"fmt"
	"runtime"
	"sync"
)

func PrintChars(name int, group *sync.WaitGroup) {
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%d: %c\n", name, ch)
		runtime.Gosched()
		// time.Sleep(time.Second)
	}
	group.Done()
}

func main() {
	group := &sync.WaitGroup{}

	n := 2

	group.Add(n)

	for i := 1; i <= n; i++ {
		// go PrintChars(i, group)
		go func(id int) {
			for ch := 'A'; ch <= 'Z'; ch++ {
				fmt.Printf("%d:%d: %c\n", id, i, ch)
				runtime.Gosched()
			}
			group.Done()
		}(i)
	}

	group.Wait()
}
