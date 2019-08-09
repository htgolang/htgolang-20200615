package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

type reqHead struct {
	method string
	url    string
	proto  string
}

func notFound(writer *bufio.Writer) {
	_, _ = writer.Write([]byte("HTTP/1.1 404 NotFound\n"))
	_, _ = writer.Write([]byte("Server: xfServer\n"))
	_, _ = writer.Write([]byte("Accept-Ranges: bytes\n"))
	_, _ = writer.Write([]byte("Content-Type: text/html\n\n"))
	_, _ = writer.Write([]byte(`<html> <head><title>404 Not Found</title></head> <body bgcolor="white"> <center><h1>404 Not Found</h1></center> </body> </html>`))
}

func Ok(writer *bufio.Writer, file []byte) {
	_, _ = writer.Write([]byte("HTTP/1.1 200 OK\n"))
	_, _ = writer.Write([]byte("Server: xfServer\n"))
	_, _ = writer.Write([]byte("Accept-Ranges: bytes\n"))
	_, _ = writer.Write([]byte("Content-Type: text/html\n\n"))
	_, _ = writer.Write(file)
}

func work(conn net.Conn) {
	var client = reqHead{}
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	bytes := make([]byte, 2048)
	n, err := reader.Read(bytes)
	if err != nil {
		if err == io.EOF {
			notFound(writer)
		}
		log.Println(n, err)
	} else {
		fmt.Printf("%s", bytes)
		request := strings.Split(string(bytes), "\n")

		client.method = strings.Split(request[0], " ")[0]
		client.url = strings.Split(request[0], " ")[1]
		client.proto = strings.Split(request[0], " ")[2]

		log.Println(client.method, client.url, client.proto)
		if stat, err := os.Stat(path + client.url); err == nil {
			if stat.IsDir() {
				if info, err := ioutil.ReadFile(path + client.url + "/index.html"); err == nil {
					Ok(writer, info)
				} else {
					notFound(writer)
				}
			} else {
				file, _ := ioutil.ReadFile(path + client.url)
				Ok(writer, file)
			}
		} else {
			notFound(writer)
		}
	}

	writer.Flush()
	conn.Close()
}

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:80")
	defer listener.Close()
	if err == nil {
		log.Println("start http success")
	} else {
		log.Println("start http failure:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		go work(conn)
	}

}
