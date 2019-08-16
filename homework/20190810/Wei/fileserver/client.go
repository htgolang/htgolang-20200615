package main

import (
	"fileserver/obj"
	"flag"
	"fmt"
	"net/rpc/jsonrpc"
	"os"
)

func main() {
	var (
		server     string
		operate    string
		path       string
		help       bool
		timeFormat string = "2006-01-02 15:04:05"
	)

	flag.StringVar(&server, "s", "127.0.0.1:9999", "Server addr")
	flag.StringVar(&operate, "o", "ls", "cmd: ls/cat")
	flag.StringVar(&path, "p", "/", "file path")
	flag.BoolVar(&help, "h", false, "help")
	flag.Usage = func() {
		fmt.Println("Usage: client [-s 127.0.0.1:9999] [-o cat/ls ] [-p /]")
		fmt.Println("Options")
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

	switch operate {
	case "ls":
		request := obj.LsRequest{Path: path}
		var response obj.LsResponse
		if err := conn.Call("Ls.Exec", &request, &response); err == nil {
			fmt.Printf("%5s %20s %10s %25s\n", "Type", "Name", "Size", "Modify Time")
			for _, fileInfo := range response.FileInfos {
				fmt.Printf("%5s %20s %10d %25s\n", fileInfo.Type, fileInfo.Name, fileInfo.Size, fileInfo.Mtime.Format(timeFormat))
			}
		} else {
			fmt.Println(err)
		}
	case "cat":
		request := obj.CatRequest{Path: path}
		var response obj.CatResponse
		if err := conn.Call("Cat.Exec", &request, &response); err == nil {
			fmt.Println(string(response.Bytes))
		} else {
			fmt.Println(err)
		}
	default:
		flag.Usage()
	}
}
