package main

import (
	"fmt"
	"net/rpc/jsonrpc"
	"rpc/objs"
)

func main() {
	// 获取连接
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:9999")
	if err == nil {

		// 定义请求对象/响应对象
		request := objs.Request{5, 12}
		var response objs.Response

		// 调用远程方法
		err := client.Call("Calc.Sum", &request, &response)

		if err == nil {
			fmt.Println(response.Result)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}
