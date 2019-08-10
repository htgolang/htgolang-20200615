package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"rpc/objs"
)

func main() {
	// 注册RPC服务
	rpc.Register(&objs.Calc{})

	server, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer server.Close()
	for {
		client, err := server.Accept()
		if err == nil {
			jsonrpc.ServeConn(client) // 使用jsonrpc处理客户端连接2
		}
	}
}
