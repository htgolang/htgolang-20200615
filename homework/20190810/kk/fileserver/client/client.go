package main

import (
	"fileserver/cmds"
	"flag"
	"fmt"
	"net/rpc/jsonrpc"
	"os"
)

func main() {
	var (
		server     string
		typ        string
		path       string
		help       bool
		dataFormat string = "2006-01-02 15:04:05"
	)
	flag.StringVar(&server, "s", "127.0.0.1:9999", "server addr")
	flag.StringVar(&typ, "t", "ls", "cmd: ls/cat")
	flag.StringVar(&path, "p", "/", "file path")
	flag.BoolVar(&help, "h", false, "help")
	flag.Usage = func() {
		fmt.Println("Usage: client [-t ls/cat] [-p /]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(0)
	}

	conn, err := jsonrpc.Dial("tcp", server)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer conn.Close()

	switch typ {
	case "ls":
		request := cmds.LsRequest{Path: path}
		var response cmds.LsResponse
		if err := conn.Call("Ls.Exec", &request, &response); err == nil {
			fmt.Printf("%4s %30s %10s %25s %25s %25s\n", "type", "name", "size", "create", "modify", "access")
			for _, fileInfo := range response.FileInfos {
				fmt.Printf("%4s %30s %10d %25s %25s %25s\n",
					fileInfo.Type, fileInfo.Name, fileInfo.Size,
					fileInfo.CreateTime.Format(dataFormat),
					fileInfo.ModifyTime.Format(dataFormat),
					fileInfo.AccessTime.Format(dataFormat),
				)
			}
		} else {
			fmt.Println(err)
		}
	case "cat":
		request := cmds.CatRequest{Path: path}
		var response cmds.CatResponse
		if err := conn.Call("Cat.Exec", &request, &response); err == nil {
			fmt.Println(string(response.Content))
		} else {
			fmt.Println(err)
		}
	default:
		flag.Usage()
	}
}
