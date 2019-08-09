package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

const (
	html404 = "404.html"
)

// 日志输出
func Log(v ...interface{}) {
	log.Println(v...)
}

// 处理连接请求
func handleConnection(conn net.Conn) {
	// 读缓存
	reader := bufio.NewReader(conn)
	// 读取请求头信息第一行，包含 请求路径
	n, err := reader.ReadString('\n')
	if err != nil {
		Log(conn.RemoteAddr().String(), " Connection Error: ", err)
		conn.Close()
	} else {
		fmt.Println("n:", n)
		// 获取url, 去除/  提取文件
		url := strings.TrimSpace(strings.Split(strings.Fields(n)[1], "/")[1])
		bb := make([]byte, 0)
		// 判断url是否存在, 如果不存在, url默认设置为404.html
		if _, err := os.Stat(url); err != nil {
			if os.IsNotExist(err) {
				url = "404.html"
			}
		}

		file, err := os.Open(url)
		bytes := make([]byte, 1024*1024)
		if err != nil {
			fmt.Println("打开文件异常:", err)
		} else {
			freader := bufio.NewReader(file)
			for {
				n, err := freader.Read(bytes)
				if err != nil {
					if err != io.EOF {
						fmt.Println(err)
					}
					break
				} else {
					// 每次读出的加入到bb []byte切片中
					for _, b := range bytes[:n] {
						bb = append(bb, b)
					}
				}
			}
		}

		// 拼接请求头
		respHeader := "HTTP/1.1 200 OK\n" +
			"Content-Type: text/html;charset=ISO-8859-1\n" +
			"Content-Length: " + string(len(bb))
		resp := respHeader + "\n\r\n" + string(bb)

		// 打印
		fmt.Println("resp:", resp)

		// 返回
		conn.Write([]byte(resp))
		conn.Close()
	}
}

func main() {
	addr := "0.0.0.0:8080"
	taddr, _ := net.ResolveTCPAddr("tcp", addr)
	listen, err := net.ListenTCP("tcp", taddr)
	if err != nil {
		fmt.Println("Listen Faliled:", addr)
		os.Exit(-1)
	}
	defer listen.Close()
	Log("等待客户端请求：")
	for {
		//接受请求连接
		conn, err := listen.AcceptTCP()
		if err != nil {
			break
		} else {
			Log(conn.RemoteAddr().String(), "接受请求.")
			//处理请求连接
			handleConnection(conn)
		}
		conn.Close()
	}
}
