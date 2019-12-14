package ens

import (
	"fmt"
	//"fmt"
	"github.com/xlotz/gocmdb/agent/config"
	"github.com/sirupsen/logrus"
	"github.com/imroc/req"
	//"time"
)

type ENS struct {
	conf *config.Config
}

func NewENS(config *config.Config) *ENS{
	return &ENS{
		conf:config,
	}
}

func (s *ENS) Start() {
	logrus.WithFields(logrus.Fields{

	}).Info("ENS 启动运行")

	go func() {
		endpoint := fmt.Sprintf("%s/heartbeat/%s/", s.conf.Endpoint, s.conf.UUID)

		for event := range s.conf.Heartbeat{
			//logrus.Debug("发送心跳消息到服务端: ", event)
			fmt.Println(event)
			response, err := req.New().Post(endpoint, req.BodyJSON(event))
			if err == nil {

				result := map[string]interface{}{}
				// 序列化
				response.ToJSON(&result)

				logrus.WithFields(logrus.Fields{
					"response": result,

				}).Debug("上传心跳信息")
			}else {
				logrus.WithFields(logrus.Fields{
					"error": err,

				}).Error("上传心跳信息失败")
			}
		}

	}()

	go func() {
		endpoint := fmt.Sprintf("%s/register/%s/", s.conf.Endpoint, s.conf.UUID)

		for event := range s.conf.Register{
			//logrus.Debug("发送心跳消息到服务端: ", event)
			fmt.Println(event)
			response, err := req.New().Post(endpoint, req.BodyJSON(event))
			if err == nil {

				result := map[string]interface{}{}
				// 序列化
				response.ToJSON(&result)

				logrus.WithFields(logrus.Fields{
					"response": result,

				}).Debug("客户端注册成功")
			}else {
				logrus.WithFields(logrus.Fields{
					"error": err,

				}).Error("客户端注册失败")
			}
		}

	}()


	go func() {
		endpoint := fmt.Sprintf("%s/log/%s/", s.conf.Endpoint, s.conf.UUID)

		for event := range s.conf.Log{
			//logrus.Debug("发送心跳消息到服务端: ", event)
			fmt.Println(event)
			response, err := req.New().Post(endpoint, req.BodyJSON(event))
			if err == nil {

				result := map[string]interface{}{}
				// 序列化
				response.ToJSON(&result)

				logrus.WithFields(logrus.Fields{
					"response": result,

				}).Debug("日志上传成功")
			}else {
				logrus.WithFields(logrus.Fields{
					"error": err,

				}).Error("日志上传失败")
			}
		}


	}()
}