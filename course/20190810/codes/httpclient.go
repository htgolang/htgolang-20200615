package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	var addr, url string

	flag.StringVar(&addr, "addr", "127.0.0.1:9999", "addr")
	flag.StringVar(&url, "url", "/net.go", "url")

	flag.Parse()

	client, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Fprintf(client, "GET %s HTTP/1.0\r\n\r\n", url)

	reader := bufio.NewReader(client)

	// 响应行
	line, err := reader.ReadString('\n')
	fmt.Println("响应行:")
	fmt.Println("\t", line)

	fmt.Println("响应头:")
	// 响应头
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		if line == "\r\n" {
			break
		}
		fmt.Print("\t", line)
	}

	fmt.Println("响应体:")

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		if line == "\r\n" {
			break
		}
		fmt.Print("\t", line)
	}

	client.Close()
}
