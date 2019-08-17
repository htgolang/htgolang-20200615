package main

import (
	"bufio"
	_ "fileserver/cmds"
	"flag"
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"os"
	"path/filepath"
)

var (
	baseDir string
)

func init() {
	path, _ := filepath.Abs(os.Args[0])
	baseDir = filepath.Dir(path)
}

func main() {
	var (
		addr    string
		help    bool
		dataDir string
	)

	flag.StringVar(&addr, "l", "127.0.0.1:9999", "listen addr")
	flag.StringVar(&dataDir, "d", filepath.Join(baseDir, "datas"), "data dir")
	flag.BoolVar(&help, "h", false, "Help")

	flag.Usage = func() {
		fmt.Println("Usage: server [-L 127.0.0.1:9999]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	logFile, err := os.OpenFile(filepath.Join(baseDir, "server.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0X066)
	if err != nil {
		fmt.Println("init log error:", err)
		os.Exit(-1)
	}

	defer logFile.Close()
	logWriter := bufio.NewWriter(logFile)
	defer logWriter.Flush()
	log.SetOutput(logWriter)

	if info, err := os.Stat(dataDir); os.IsNotExist(err) {
		os.MkdirAll(dataDir, 0X066)
		log.Print("[info] init data dir: ", dataDir)
	} else if !info.IsDir() {
		log.Fatal("[error] data dir is not directory: ", dataDir)
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Print("[error] fileserver start error: ", err)
		os.Exit(-1)
	}
	defer listener.Close()
	log.Print("[info] fileserver listen on: ", listener.Addr())
	log.Print("[info] fileserver data dir: ", dataDir)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print("[error] client accept error: ", err)
			continue
		}
		log.Print("[info] client is connected: ", conn.RemoteAddr())
		go jsonrpc.ServeConn(conn)
	}
}
