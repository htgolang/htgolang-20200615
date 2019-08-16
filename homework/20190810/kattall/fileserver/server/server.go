package main

import (
	"flag"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"path/filepath"
	"fileserver/cmds"
)

var (
	basedir string
)

func init(){
	path, _ := filepath.Abs(os.Args[0])
	basedir = filepath.Dir(path)
}

func main() {
	var (
		addr    string
		help    bool
		datadir string
	)
	fmt.Println("basedir", basedir)
	flag.StringVar(&addr, "l", "0.0.0.0:8888", "listen addr")
	flag.StringVar(&datadir, "d","fileserver","fileserver dir")
	flag.BoolVar(&help, "h", false, "Help")

	flag.Usage = func() {
		fmt.Println("Usage: server.exe [-l 0.0.0.0:8888]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if info, err := os.Stat(datadir); os.IsNotExist(err) {
		os.MkdirAll(datadir, os.ModePerm)
		fmt.Println("[info] init data dir:", datadir)
	} else if !info.IsDir() {
		fmt.Println("[error] data dir is not directory:", datadir)
	}

	rpc.Register(&cmds.Ls{Basedir:datadir})
	fmt.Println("[info] Rpc register cmds.Ls")

	rpc.Register(&cmds.Cat{Basedir:datadir})
	fmt.Println("[info] Rpc register cmds.Cat")

	rpc.Register(&cmds.Delete{Basedir:datadir})
	fmt.Println("[info] Rpc register cmds.Delete")

	rpc.Register(&cmds.Put{Basedir:datadir})
	fmt.Println("[info] Rpc register cmds.Put")

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("[error] Listen %s failed:%s", addr, err)
		os.Exit(-1)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("%s client accept failed: %s", conn, err)
			continue
		}
		fmt.Println("[info] Client Accept connected:", conn.RemoteAddr())
		go jsonrpc.ServeConn(conn)
	}

}