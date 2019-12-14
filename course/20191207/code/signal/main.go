package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	go func() {
		for now := range time.Tick(time.Second) {
			fmt.Println(now)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("wait")
	<-ch
	fmt.Println("over")
}
