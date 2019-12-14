package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/imsilence/gocmdb/agent/config"
	"github.com/imsilence/gocmdb/agent/ens"

	"github.com/imsilence/gocmdb/agent/plugins"
	_ "github.com/imsilence/gocmdb/agent/plugins/init"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	gconf, err := config.NewConfig()
	if err != nil {
		logrus.Error("读取配置出错:", err)
		os.Exit(-1)
	}
	defer func() {
		os.Remove(gconf.PidFile)
	}()
	log, err := os.OpenFile(gconf.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		logrus.Error("打开日志文件出错:", err)
		os.Exit(-1)
	}
	defer func() {
		log.Close()
	}()

	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{})
	// logrus.SetOutput(log)
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
