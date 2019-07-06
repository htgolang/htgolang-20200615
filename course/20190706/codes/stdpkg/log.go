package main

import "log"

func main() {

	log.Printf("我是Printf日志: %s", "x")

	log.SetPrefix("prefix:")
	log.SetFlags(log.Flags() | log.Llongfile)

	log.Printf("我是Printf日志: %s", "x")

	// log.Panicf("我是Panic日志: %s", "y")
	log.Fatalf("我是Fatalf日志: %s", "z")
	log.Fatalf("我是Fatalf日志: %s", "z")
}
