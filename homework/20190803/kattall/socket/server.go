package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	addr := "0.0.0.0:9999"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer listener.Close()
	fmt.Println("Listen: ", addr)
	input := bufio.NewScanner(os.Stdin)

	for {
		client, err := listener.Accept()
		if err == nil {
			reader := bufio.NewReader(client)
			writer := bufio.NewWriter(client)
			fmt.Printf("客户端%s连接成功\n", client.RemoteAddr())
			for {
				fmt.Print("请输入(q退出):")
				input.Scan()
				if input.Text() == "q" {
					break
				}
				_, err := writer.WriteString(input.Text() + "\n")
				writer.Flush()
				if err != nil {
					break
				}

				line, err := reader.ReadString('\n')
				if err != nil {
					break
				}
				fmt.Println("客户端:", strings.TrimSpace(line))
			}
			fmt.Printf("客户端%s关闭\n", client.RemoteAddr())
			client.Close()
		}
	}
}
