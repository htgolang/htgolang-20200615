package ens

import (
	"fmt"
	"github.com/dcosapp/gocmdb/agent/config"
	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
)

type ENS struct {
	conf *config.Config
}

func NewENS(conf *config.Config) *ENS {
	return &ENS{conf: conf}
}

func (s *ENS) Start() {
	//req.Debug = true
	logrus.Info("ENS 开始运行")
	go func() {
		endpoint := fmt.Sprintf("%s/heartbeat/%s/", s.conf.Endpoint, s.conf.UUID)
		// 从chan读数据
		for evt := range s.conf.Heartbeat {
			// 发送数据至server
			response, err := req.New().Post(endpoint, req.BodyJSON(evt))
			if err == nil {
				result := map[string]interface{}{}
				response.ToJSON(&result)
				logrus.WithFields(logrus.Fields{"response": result,}).Debug("上传心跳消息")
			} else {
				logrus.WithFields(logrus.Fields{"error": err.Error(),}).Error("上传心跳消息失败")
			}
		}
	}()

	go func() {
		endpoint := fmt.Sprintf("%s/register/%s/", s.conf.Endpoint, s.conf.UUID)
		// 从chan读数据
		for evt := range s.conf.Register {
			// 发送数据至server
			response, err := req.New().Post(endpoint, req.BodyJSON(evt))
			if err == nil {
				result := map[string]interface{}{}
				response.ToJSON(&result)
				logrus.WithFields(logrus.Fields{"response": result,}).Debug("注册成功")
			} else {
				logrus.WithFields(logrus.Fields{"error": err.Error(),}).Error("注册失败")
			}
		}
	}()

	go func() {
		endpoint := fmt.Sprintf("%s/log/%s/", s.conf.Endpoint, s.conf.UUID)
		// 从chan读数据
		for evt := range s.conf.Log {
			// 发送数据至server
			response, err := req.New().Post(endpoint, req.BodyJSON(evt))
			if err == nil {
				result := map[string]interface{}{}
				response.ToJSON(&result)
				logrus.WithFields(logrus.Fields{"response": result,}).Debug("日志上传成功")
			} else {
				logrus.WithFields(logrus.Fields{"error": err.Error(),}).Error("日志上传失败")
			}
		}
	}()
}
