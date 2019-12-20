package main

import (
	"os"
	"os/signal"
	"syscall"
	"flag"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/imroc/req"

	"github.com/imsilence/gocmdb/agent/config"
	"github.com/imsilence/gocmdb/agent/ens"

	"github.com/imsilence/gocmdb/agent/plugins"
	_ "github.com/imsilence/gocmdb/agent/plugins/init"
)

func main() {
	h := flag.Bool("h", false, "help")
	help := flag.Bool("help", false, "help")
	verbose := flag.Bool("v", false, "verbose")
	flag.Usage = func() {
		fmt.Println("agentd")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *h || *help {
		flag.Usage()
		os.Exit(0)
	}

	reader := viper.New()
	reader.SetConfigName("agent")
	reader.AddConfigPath("etc")
	reader.SetConfigType("yaml")
	if err := reader.ReadInConfig(); err != nil {
		logrus.Error("读取配置出错:", err)
		os.Exit(-1)
	}

	gconf, err := config.NewConfig(reader)
	if err != nil {
		logrus.Error("读取配置出错:", err)
		os.Exit(-1)
	}

	defer func() {
		os.Remove(gconf.PidFile)
	}()

	if *verbose {
		logrus.SetLevel(logrus.DebugLevel)
		req.Debug = true
	} else {
		logrus.SetLevel(logrus.InfoLevel)
		log, err := os.OpenFile(gconf.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			logrus.Error("打开日志文件出错:", err)
			os.Exit(-1)
		}
		defer func() {
			log.Close()
		}()
		logrus.SetOutput(log)
	}

	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{})

	logrus.WithFields(logrus.Fields{
		"PID":  gconf.PID,
		"UUID": gconf.UUID,
	}).Info("Agent启动")

	plugins.DefaultManager.Init(gconf)

	ens.NewENS(gconf).Start()
	plugins.DefaultManager.Start()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	logrus.WithFields(logrus.Fields{
		"PID":  gconf.PID,
		"UUID": gconf.UUID,
	}).Info("Agent退出")
}
