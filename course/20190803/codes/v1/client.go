package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr := "127.0.0.1:9999"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	bytes := make([]byte, 1024)
	if n, err := conn.Read(bytes); err == nil {
		fmt.Println(string(bytes[:n]))
	} else {
		fmt.Println(err)
	}
	conn.Close()
}
