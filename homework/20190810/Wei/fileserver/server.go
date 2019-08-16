package main

import (
	"bufio"
	"fileserver/obj"
	"flag"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"path/filepath"
)

var baseDir string

func init() {
	path, _ := filepath.Abs(os.Args[0])
	baseDir = filepath.Dir(path)
}

func main() {
	var (
		addr    string
		dataDir string
		help    bool
	)

	flag.StringVar(&addr, "l", "127.0.0.1:9999", "listen addr")
	flag.StringVar(&dataDir, "d", filepath.Join(baseDir, "data"), "data dir")
	flag.BoolVar(&help, "h", false, "help")
	flag.Usage = func() {
		fmt.Println("Usage: server [-l 127.0.0.1:9999] [-d data]")
		fmt.Println("Options")
		flag.PrintDefaults()
	}
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	logPath := filepath.Join(baseDir, "server.log")
	log_file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err == nil {
		log.SetOutput(log_file)
		commandfile := os.Args[0]
		log.SetPrefix(commandfile + ":")
		log.SetFlags(log.Flags() | log.Lshortfile)
	} else {
		fmt.Println("init log error: ", err)
		os.Exit(-1)
	}

	defer log_file.Close()

	logWrite := bufio.NewWriter(log_file)
	defer logWrite.Flush()

	if info, err := os.Stat(dataDir); os.IsNotExist(err) {
		os.MkdirAll(dataDir, 0X066)
		log.Print("[info] init data dir: ", dataDir)
	} else if !info.IsDir() {
		log.Fatal("[error] data dir is not directory: ", dataDir)
	}

	rpc.Register(&obj.Ls{BaseDir: dataDir})
	log.Print("[info] register obj.Ls")

	rpc.Register(&obj.Cat{BaseDir: dataDir})
	log.Print("[info] register obj.Cat")

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Print("[error] fileserver listen error: ", err)
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

		log.Print("[info] client is connection: ", conn.RemoteAddr())
		go jsonrpc.ServeConn(conn)
	}
}
