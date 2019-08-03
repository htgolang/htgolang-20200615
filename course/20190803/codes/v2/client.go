package main

import (
	"bufio"
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
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	input := bufio.NewScanner(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print("服务器:", line)
		fmt.Print("请输入(q退出):")
		input.Scan()
		if input.Text() == "q" {
			break
		}
		_, err = writer.WriteString(input.Text() + "\n")
		writer.Flush()
		if err != nil {
			break
		}

	}

}
