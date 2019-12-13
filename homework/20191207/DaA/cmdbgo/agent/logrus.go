package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main(){
	//日志等级
	logrus.SetLevel(logrus.DebugLevel)
	//记录日志
	logrus.Info("info")
	logrus.Debug("debug")
	logrus.Warn("warn")
	logrus.Error("error")

	//占位符使用
	logrus.Debugf("我是占位%s","DaA")

	//调用关系记录
	logrus.SetReportCaller(true)
	logrus.Error("记录下调用关系")

	//常规日志格式
	logrus.WithFields(logrus.Fields{
		"module": "main",
		"test": "xxx",
	}).Info("常规日志")

	//json格式日志
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.WithFields(logrus.Fields{
		"module": "main",
		"test": "xxx",
	}).Info("我是json格式的日志")

	//记录到文件
	logfile, _ := os.OpenFile("main.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0777)
	defer func(){
		logfile.Close()
	}()
	logrus.SetOutput(logfile)
	logrus.Info("日志记录到文件")

	//占位符使用
	logrus.Infof("我是info占位%s","DaA")

}
