package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// 存放所有client的写缓冲
var clientInfo = make(map[string]*bufio.Writer)

func message(client *net.TCPConn) {
	// 生成读缓冲
	reader := bufio.NewReader(client)
	for {
		// 读到client写入的内容
		readString, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		// 循环将内容发送给除了自己以外的所有client
		for key, conn := range clientInfo {
			if key == client.RemoteAddr().String() {
				continue
			}

			_, err = conn.WriteString(readString)
			if err != nil {
				fmt.Println("write failure")
				break
			}

			err = conn.Flush()
			if err != nil {
				fmt.Println("fLush failure")
				break
			}
		}

	}
}

func main() {
	// 监听本地8080端口
	addr := "0.0.0.0:9999"
	tcpAddr, _ := net.ResolveTCPAddr("tcp", addr)
	conn, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Printf("Listen%s: %s\n", addr, err)
		os.Exit(-1)
	}

	defer conn.Close()
	log.Println("Listen:", addr)

	for {
		// 等待客户端连接
		client, err := conn.AcceptTCP()
		if err == nil {
			defer client.Close()

			// 生成写缓冲
			writer := bufio.NewWriter(client)

			// 将写缓冲放入clientInfo中
			clientInfo[client.RemoteAddr().String()] = writer
			log.Printf("远程客户端: %s 上线了!\n", client.RemoteAddr().String())

			// 将每个client都扔给goroutine去异步执行
			go message(client)
		}
	}
}
