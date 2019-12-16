package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	logrus.Debug("我是一个debug信息")
	logrus.Info("我是一个info信息")
	logrus.Warn("我是一个warn信息")
	logrus.Error("我是一个error信息")

	logrus.WithFields(logrus.Fields{
		"module": "main",
		"test":   "xxx",
	}).Info("我是一个带fields的日志")

	logrus.SetReportCaller(true)
	logrus.Error("我记录的调用关系")

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.WithFields(logrus.Fields{
		"module": "main",
		"test":   "xxx",
	}).Info("我是一个json格式的日志")

	logfile, _ := os.OpenFile("main.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	defer func() {
		logfile.Close()
	}()
	logrus.SetOutput(logfile)

	logrus.Info("我记录日志到文件")

	log := logrus.New()

	log.SetLevel(logrus.DebugLevel)
	log.Debug("tet")
	log.Debug("我是:", "kk")
}
