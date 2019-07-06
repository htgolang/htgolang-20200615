package main

import (
	"flag"
	"fmt"
)

func main() {

	port := flag.Int("P", 22, "ssh port")
	host := flag.String("H", "127.0.0.1", "ssh host")
	verbor := flag.Bool("v", false, "detail log")
	help := flag.Bool("h", false, "help")

	flag.Usage = func() {
		fmt.Println("usage: flagargs [-H 127.0.0.14] [-P 22] [-v] [args01 args02]")
		flag.PrintDefaults()
	}

	// 解析命令行参数
	flag.Parse()

	fmt.Printf("%T %T, %T , %T\n", port, host, verbor, help)
	fmt.Printf("%T %T, %T , %T\n", *port, *host, *verbor, *help)
	fmt.Printf("%v %v, %v , %v\n", *port, *host, *verbor, *help)

}
