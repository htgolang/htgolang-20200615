package main

import (
	"flag"
	"fmt"
	"github.com/dcosapp/gocmdb/agent/config"
	"github.com/dcosapp/gocmdb/agent/ens"
	"github.com/dcosapp/gocmdb/agent/plugins"
	_ "github.com/dcosapp/gocmdb/agent/plugins/init"
	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	verbose := flag.Bool("v", false, "verbose")
	flag.Usage = func() {
		fmt.Println("usage: agentd -h")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *verbose {
		// 日志级别为Debug
		logrus.SetLevel(logrus.DebugLevel)
		req.Debug = true
	} else {
		logrus.SetLevel(logrus.InfoLevel)
		req.Debug = false
	}

	configReader := viper.New()
	configReader.SetConfigName("agent")
	configReader.SetConfigType("yaml")
	configReader.AddConfigPath("etc/")
	err := configReader.ReadInConfig()
	if err != nil {
		logrus.Error("读取配置出错:", err)
		os.Exit(-1)
	}

	// 初始化程序配置
	gconf, err := config.NewConfig(configReader)
	if err != nil {
		logrus.Error("读取配置出错:", err)
		os.Exit(-1)
	}
	defer func() {
		_ = os.Remove(gconf.PidFile)
	}()

	// 打开或生成日志文件
	log, err := os.OpenFile(gconf.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		logrus.Error("打开日志文件出错:", err)
		os.Exit(-1)
	}
	defer func() {
		log.Close()
	}()

	// 日志为文本格式
	logrus.SetFormatter(&logrus.TextFormatter{})

	// 设置日志输出到日志文件
	logrus.SetOutput(log)

	logrus.WithFields(logrus.Fields{"PID": gconf.PID, "UUID": gconf.UUID,}).Info("Agent启动")

	// 初始化插件
	plugins.DefaultManager.Init(gconf)

	// 启动ens上传至server段
	ens.NewENS(gconf).Start()

	// 启动其他插件将数据上传给ens
	plugins.DefaultManager.Start()

	// 阻塞监听ctrl+c信号
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	//
	<-ch
	logrus.WithFields(logrus.Fields{
		"PID":  gconf.PID,
		"UUID": gconf.UUID,
	}).Info("Agent退出")

}
