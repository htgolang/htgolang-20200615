package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/JevonWei/gocmdb/agent/config"
	"github.com/JevonWei/gocmdb/agent/ens"
	"github.com/JevonWei/gocmdb/agent/plugins"
	_ "github.com/JevonWei/gocmdb/agent/plugins/init"
	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
)

func main() {
	// 初始化命令行参数
	h := flag.Bool("h", false, "help")
	help := flag.Bool("help", false, "help")
	verbose := flag.Bool("v", false, "verbose")

	flag.Usage = func() {
		fmt.Println("usage: agentd -h")
		flag.PrintDefaults()
	}
	// 解析命令行参数
	flag.Parse()

	if *h || *help {
		flag.Usage()
		os.Exit(0)
	}

	if !*verbose {
		//设置日志级别为Info
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
		req.Debug = true
	}

	gconf, err := config.NewConfig()
	if err != nil {
		logrus.Error("读取配置出错: ", err)
		os.Exit(-1)
	}

	defer func() {
		os.Remove(gconf.PidFile)
	}()

	log, err := os.OpenFile("./logs/agent.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		logrus.Error("打开日志文件出错:", err)
		os.Exit(-1)
	}
	defer func() {
		log.Close()
	}()

	// 将日志信息输出到终端
	logrus.SetFormatter(&logrus.TextFormatter{})
	//logrus.SetFormatter(&logrus.JSONFormatter{})

	// 将日志信息输出到log文件
	//logrus.SetOutput(log)

	logrus.WithFields(logrus.Fields{
		"PID":  gconf.PID,
		"UUID": gconf.UUID,
	}).Info("Agent启动")

	plugins.DefaultManager.Init(gconf)
	ens.NewENS(gconf).Start()

	plugins.DefaultManager.Start()
	// go func() {
	// 	for now := range time.Tick(time.Second) {
	// 		logrus.Info(now)
	// 	}
	// }()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	logrus.WithFields(logrus.Fields{
		"PID":  gconf.PID,
		"UUID": gconf.UUID,
	}).Info("Agent退出")

}
