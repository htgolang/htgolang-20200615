package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	dirPath string
	defaultFileName string
	file404 string
	file500 string
)

func handle(conn net.Conn) {
	defer func() {
		conn.Close()
		fmt.Printf("client %s close\n", conn.RemoteAddr())
	}()
	fmt.Printf("来自 %s 连接成功\n ", conn.RemoteAddr())

	time.Sleep(time.Second * 10)

	// 处理客户端请求

	reader := bufio.NewReader(conn)

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		//break
	} else {
		// 处理正常请求
		nodes := strings.Fields(line)

        // 获取请求资源路径
		path := filepath.Join(dirPath, nodes[1])

		if info, err := os.Stat(path); err != nil {
			if os.IsNotExist(err) {
				path = file404
			} else {
				path = file500
			}
		} else {
			//目录
			if info.IsDir() {
				path = filepath.Join(path, defaultFileName)
			}
			//文件
		}
		// 对文件path验证

		if _, err := os.Stat(path); err == nil {
			fmt.Fprint(conn, "HTTP/1.1 200 ok \r\n")
			fmt.Fprint(conn, "Server: httpserver\r\n")
			fmt.Fprint(conn, "\r\n")

			bytes, _ := ioutil.ReadFile(path)
			conn.Write(bytes)
		} else {
			fmt.Fprint(conn, "HTTP/1.1 404 NotFount\r\n")
		}

	}

}

func init() {
	binPath, _ := filepath.Abs(os.Args[0])

	dirPath = filepath.Dir(binPath)

	defaultFileName = "index.html"

	file404 = filepath.Join(dirPath, "40x.html")
	file500 = filepath.Join(dirPath, "50x.html")

}


func main()  {
	addr := ":19999"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("Listen: ", addr)
	defer listener.Close()


	for {
	    client, err := listener.Accept()
	    if err != nil {
			fmt.Println(err)
			continue
		}
	    go handle(client)

	}

}
