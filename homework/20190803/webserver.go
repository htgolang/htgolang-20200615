package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	baseDir string
	timeout time.Duration
)

type Request struct {
	Method   string
	URL      string
	Protocol string
	Headers  map[string]string
	Body     []byte
}

func NewRequest(c net.Conn) (*Request, error) {
	reader := bufio.NewReader(c)
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	nodes := strings.Fields(line)
	method, url, protocol := nodes[0], nodes[1], nodes[2]
	headers := make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if line == "\r\n" {
			break
		}
		nodes := strings.SplitN(line, ":", 2)
		headers[strings.TrimSpace(nodes[0])] = strings.TrimSpace(nodes[1])
	}
	bytes := make([]byte, 1024)
	for {
		n, err := reader.Read(bytes)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(bytes[:n]))
	}

	return &Request{
		Method:   method,
		URL:      url,
		Protocol: protocol,
		Headers:  headers,
	}, nil
}

func (r *Request) Handle() *Response {
	var response *Response
	path := filepath.Join(baseDir, r.URL)

	if info, err := os.Stat(path); err == nil && info.IsDir() {
		path = filepath.Join(path, "index.html")
	}

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			path = filepath.Join(baseDir, "404.html")
		}
	}

	bytes, err := ioutil.ReadFile(path)
	if err == nil {
		response = NewResponse(200, "OK", bytes)
	} else if os.IsNotExist(err) {
		response = NewResponse(404, "Not Found", nil)
	} else {
		response = NewResponse(500, "Internal Server Error", nil)
	}

	return response
}

type Response struct {
	Protocol   string
	Status     int
	StatusText string
	Headers    map[string]string
	Content    []byte
	Buffer     *bytes.Buffer
}

func NewResponse(status int, text string, content []byte) *Response {
	protocol := "HTTP/1.0"
	headers := map[string]string{
		"Server":       "WebServer1.0",
		"Content-Type": "text/html",
	}
	buffer := bytes.NewBufferString(fmt.Sprintf("%s %d %s\r\n", protocol, status, text))
	for k, v := range headers {
		buffer.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	buffer.WriteString("\r\n")
	buffer.Write(content)
	return &Response{
		Protocol:   protocol,
		Status:     status,
		StatusText: text,
		Headers:    headers,
		Content:    content,
		Buffer:     buffer,
	}
}

func (r *Response) Read(bytes []byte) (int, error) {
	return r.Buffer.Read(bytes)
}

func init() {
	filePath, _ := filepath.Abs(os.Args[0])
	baseDir = filepath.Dir(filePath)

	timeout = time.Second * 2
}

func main() {

	addr := ":9999"
	server, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer server.Close()

	for {
		client, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go func(c net.Conn) {
			defer c.Close()
			request, err := NewRequest(c)
			var response *Response
			if err == nil {
				response = request.Handle()
			} else {
				response = NewResponse(500, "Internal Server Error", nil)
			}
			io.Copy(c, response)
		}(client)

	}
}
