package main

import (
	"flag"
	"fmt"
	"github.com/xlotz/rpcs/objects"
	"net/rpc/jsonrpc"
	"os"
)

func main()  {
	var (
		server string
		typ string
		path string
		help bool
		dataFormat string = "2006-01-02 15:04:05"
	)

	flag.StringVar(&server, "s", "127.0.0.1:10000", "server addr")
	flag.StringVar(&typ, "t","ls", "cmd: ls/cat")
	flag.StringVar(&path, "p", "/", "file path")
	flag.BoolVar(&help, "h", false, "help")
	flag.Usage = func() {
		fmt.Println("Usage: client [-t ls/cat] [-p /]")
		fmt.Println("Options: ")
		flag.PrintDefaults()
	}

	flag.Parse()

	if help{
		flag.Usage()
		os.Exit(0)
	}




	// 获取连接

	client, err := jsonrpc.Dial("tcp", server)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer client.Close()

	switch typ {
	case "ls":
		request := objects.LsRequest{Path: path}
		var response objects.LsResponse

		if err := client.Call("Ls.Exec", &request, &response); err == nil {
			fmt.Printf("%3s %30s %10s %25s\n", "type", "name", "size", "create")
			for _, fileinfo := range response.FileInfos {
				fmt.Printf("%4s %30s %10d %25s\n", fileinfo.Type, fileinfo.Name, fileinfo.Size,
					fileinfo.ModifyTime.Format(dataFormat),)
			}
		}else {
			fmt.Println(err)
		}
	case "cat":
		request := objects.CatRequest{Path: path}
		var response objects.CatResponse
		if err := client.Call("Cat.Exec", &request, &response); err == nil {
			fmt.Println(string(response.Content))
		}else {
			fmt.Println(err)
		}
	default:
		flag.Usage()

	}
}
