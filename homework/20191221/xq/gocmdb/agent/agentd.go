package main

import (
	"github.com/sirupsen/logrus"
	"github.com/xlotz/gocmdb/agent/config"
	"github.com/xlotz/gocmdb/agent/plugins"
	_ "github.com/xlotz/gocmdb/agent/plugins/init"
	"github.com/xlotz/gocmdb/agent/ens"
	"github.com/spf13/viper"

	"os"
	"os/signal"
	"syscall"
	//"time"
)

func main()  {

	logrus.SetLevel(logrus.DebugLevel)

	configReader := viper.New()
	configReader.SetConfigName("agent")
	configReader.SetConfigType("yaml")
	configReader.AddConfigPath("etc/")

	err := configReader.ReadInConfig()
	if err != nil {
		logrus.Error("读取配置出错", err)
		os.Exit(-1)
	}


	//gconf, err := config.NewConfig()
	gconf, err := config.NewConfig(configReader)
	if err != nil {
		logrus.Error("读取配置出错", err)
		os.Exit(-1)
	}

	defer func(){
		os.Remove(gconf.PIDFile)
	}()
	//返回file 指针
	logfile, err:= os.OpenFile(gconf.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil {
		logrus.Error("打开日志文件出错: ",err)
		os.Exit(-1)
	}
	defer func() {
		logfile.Close()
	}()

	logrus.SetFormatter(&logrus.TextFormatter{})
	//logrus.SetFormatter(&logrus.JSONFormatter{})

	//logrus.SetOutput(logfile)

	logrus.WithFields(logrus.Fields{
		"PID": gconf.PID,
		"UUID": gconf.UUID,
	}).Info("Agent 启动")

	plugins.DefaultManager.Init(gconf)

	ens.NewENS(gconf).Start()


	plugins.DefaultManager.Start()


	//go func() {
	//	for now := range time.Tick(time.Second){
	//		logrus.Info(now)
	//	}
	//
	//}()

	//创建历程
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	<-ch

	logrus.WithFields(logrus.Fields{
		"PID": gconf.PID,
		"UUID": gconf.UUID,
	}).Info("Agent 退出")




}
