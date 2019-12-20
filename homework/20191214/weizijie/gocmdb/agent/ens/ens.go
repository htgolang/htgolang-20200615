package ens

import (
	"fmt"

	"github.com/JevonWei/gocmdb/agent/config"
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
	// 打印Debug信息，将响应信息打印
	// req.Debug = true

	logrus.Info("ENS 开始执行")
	go func() {
		//endpoint := s.conf.Endpoint + "/heartbeat/" + s.conf.UUID + "/"
		endpoint := fmt.Sprintf("%s/heartbeat/%s/", s.conf.Endpoint, s.conf.UUID)
		for evt := range s.conf.Heartbeat {
			response, err := req.New().Post(endpoint, req.BodyJSON(evt))
			if err == nil {
				// 将response反序列化为JSON
				result := map[string]interface{}{}
				response.ToJSON(&result)

				logrus.WithFields(logrus.Fields{
					"result": result,
					// "response": response,
				}).Debug("上传心跳信息成功")
			} else {
				logrus.WithFields(logrus.Fields{
					"error": err,
				}).Debug("上传心跳信息失败")
			}
		}
	}()

	go func() {
		//endpoint := s.conf.Endpoint + "/heartbeat/" + s.conf.UUID + "/"
		endpoint := fmt.Sprintf("%s/register/%s/", s.conf.Endpoint, s.conf.UUID)
		for evt := range s.conf.Register {
			response, err := req.New().Post(endpoint, req.BodyJSON(evt))
			if err == nil {
				// 将response反序列化为JSON
				result := map[string]interface{}{}
				response.ToJSON(&result)

				logrus.WithFields(logrus.Fields{
					"result": result,
					// "response": response,
				}).Debug("注册成功")
			} else {
				logrus.WithFields(logrus.Fields{
					"error": err,
				}).Debug("注册失败")
			}
		}
	}()

	go func() {
		//endpoint := s.conf.Endpoint + "/heartbeat/" + s.conf.UUID + "/"
		endpoint := fmt.Sprintf("%s/log/%s/", s.conf.Endpoint, s.conf.UUID)
		for evt := range s.conf.Log {
			response, err := req.New().Post(endpoint, req.BodyJSON(evt))
			if err == nil {
				// 将response反序列化为JSON
				result := map[string]interface{}{}
				response.ToJSON(&result)

				logrus.WithFields(logrus.Fields{
					"result": result,
					// "response": response,
				}).Debug("日志上传成功")
			} else {
				logrus.WithFields(logrus.Fields{
					"error": err,
				}).Debug("日志上传失败")
			}
		}
	}()
}
