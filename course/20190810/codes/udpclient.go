package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	addr := "127.0.0.1:9999"
	conn, err := net.Dial("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "Time: %d", time.Now().Unix())

	bytes := make([]byte, 1024)
	n, err := conn.Read(bytes)
	fmt.Println(n, err, string(bytes[:n]))
}
