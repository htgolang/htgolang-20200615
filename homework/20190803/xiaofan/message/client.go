package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func read(reader *bufio.Reader, name string) {
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		// 格式化输出
		fmt.Println()
		log.Println(readString)
		fmt.Print(name)
	}
}
func main() {
	// 连接server
	addr := "127.0.0.1:9999"
	tcpAddr, _ := net.ResolveTCPAddr("tcp", addr)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	input := bufio.NewScanner(os.Stdin)

	// 定义昵称
	fmt.Print("在加入聊天室前，请输入你的昵称：")
	input.Scan()
	name := "<" + input.Text() + "(你)>:"
	nickname := "<" + input.Text() + ">:"

	// 将缓冲读送到goroutine异步执行
	go read(reader, name)

	// 主goroutine负责将输入的内容写给server
	for {
		fmt.Print(name)
		input.Scan()
		if input.Text() == "q" {
			break
		}

		if input.Text() == "" {
			break
		}

		trimStr := strings.TrimSpace(input.Text())
		_, err = writer.WriteString(nickname + trimStr + "\n")
		if err != nil {
			fmt.Println(err)
			break
		}
		_ = writer.Flush()
	}
}
