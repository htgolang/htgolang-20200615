package main

import (
	"fileserver/cmds"
	"flag"
	"fmt"
	"io/ioutil"
	"net/rpc/jsonrpc"
	"os"
)

func main(){
	var (
		server string
		typ string
		path string
		help bool
		src string
		dataFormat string = "2006-01-02 15:04:05"
	)

	flag.StringVar(&server, "s","127.0.0.1:8888","server addr")
	flag.StringVar(&typ, "t","ls","cmd: ls/cat")
	flag.StringVar(&path, "p","/","file path")
	flag.StringVar(&src, "src", "", "src path")
	flag.BoolVar(&help, "h", false, "Help")

	flag.Usage = func() {
		fmt.Println("Usage: client.exe [-s 0.0.0.0:8888] [-t ls/cat/put] [-p /]")
		fmt.Println("ls/cat/delete command usage: \n\t-s specifies server address and port \n\t-t specifies the command -p specifies path")
		fmt.Println("Usage: client.ext [-s 0.0.0.0:8888] [-t put] [-p /] [-src srcpath]")
		fmt.Println("put command usage: \n\t-s specifies server address and port  \n\t-t specifies the command \n\t-p specifies upload path \n\t-src specifies local path")
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
		fmt.Printf("[error] Connection %s server failed: %s", server, err)
		os.Exit(-1)
	}
	defer conn.Close()

	switch typ {
	case "ls":
		reqeuest := cmds.LsRequest{Path:path}
		var response cmds.LsResponse
		if err := conn.Call("Ls.Exec", &reqeuest, &response); err == nil {
			fmt.Printf("%4s %30s %10s %25s\n", "Type", "Name", "Size", "ModifyTime")
			for _, fileinfo := range response.FileInfos {
				fmt.Printf("%4s %30s %10d %25s\n", fileinfo.Type, fileinfo.Name, fileinfo.Size, fileinfo.ModifyTime.Format(dataFormat))
			}
		} else {
			fmt.Println("发生错误：", err)
		}
	case "cat":
		request := cmds.CatRequest{Path:path}
		var response cmds.CatResponse
		if err := conn.Call("Cat.Exec", &request, &response); err == nil {
			fmt.Println(string(response.Content))
		} else {
			fmt.Printf("发生错误：%s", err)
		}
	case "delete":
		request := cmds.DeleteRequest{Path:path}
		var response cmds.DeleteResponse
		if err := conn.Call("Delete.Exec", &request, &response); err == nil {
			fmt.Println(response.Message)
		} else {
			fmt.Printf("发生错误: %s", err)
		}
	case "put":
		request := cmds.PutRequest{Path:path, SrcPath: src}
		var response cmds.PutResponse

		if file, err := os.Stat(request.SrcPath); err != nil {
			if os.IsNotExist(err) {
				fmt.Println("源文件不存在.")
				return
			}
			return
		} else {
			if file.IsDir() {
				fmt.Println("不能上传目录")
				return
			}
		}

		bytes, err := ioutil.ReadFile(request.SrcPath)
		if err == nil {
			request.Content = bytes
			//fmt.Println("bytes: ", string(bytes))
		} else {
			fmt.Println("文件读取失败")
			return
		}

		if err := conn.Call("Put.Exec", &request, &response); err == nil {
			fmt.Println(response.Message)
		} else {
			fmt.Println(err)
		}
	default:
		flag.Usage()
	}
}
