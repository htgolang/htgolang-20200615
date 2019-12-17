package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xxdu521/cmdbgo/agent/config"
	"github.com/xxdu521/cmdbgo/agent/ens"
	"github.com/xxdu521/cmdbgo/agent/plugins"
	"os"
	"os/signal"
	"syscall"
	_ "github.com/xxdu521/cmdbgo/agent/plugins/init"
)

func main(){
	varbose := flag.Bool("v",false,"debug")
	flag.Usage = func(){
		fmt.Println("usage: agent -h")
		flag.PrintDefaults()
	}
	flag.Parse()

	//读取配置信息
	gconf,err := config.NewConfig()
	if err != nil {
		logrus.Error("读取配置出错")
		os.Exit(-1)
	}
	defer func(){
		os.Remove(gconf.PidFile)
	}()

	//打开日志文件
	logfile, err := os.OpenFile("logs/agent.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		logrus.Error("日志文件打开失败", err)
		os.Exit(-1)
	}
	defer func(){ logfile.Close() }()

	//日志级别设置
	if !*varbose {
		logrus.SetLevel(logrus.DebugLevel)  //设置info日志级别
		logrus.SetOutput(logfile) //输入日志到文件
	} else {
		logrus.SetLevel(logrus.DebugLevel)  //设置debug日志级别
		//日志默认打印在控制台
	}
	//logrus.SetFormatter(&logrus.TextFormatter{}) //文本格式
	logrus.SetFormatter(&logrus.JSONFormatter{}) //Json格式


	logrus.WithFields(logrus.Fields{
		"PID": gconf.PID,
		"UUID": gconf.UUID,
	}).Info("agent启动成功")

	//加载配置，agent启动成功之后，调用管理器的初始化方法，进行插件的初始化工作
	plugins.DefaultManager.Init(gconf)

	//初始化ENS，并启动ENS服务
	ens.NewENS(gconf).Start()
	//启动插件方法
	plugins.DefaultManager.Start()

	/* 每秒钟打印一次数据
	go func(){
		for now := range time.Tick(time.Second * 1){
			logrus.Info(now)
		}
	}()
	 */

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	logrus.WithFields(logrus.Fields{
		"PID": gconf.PID,
		"UUID": gconf.UUID,
	}).Info("agent退出")

}