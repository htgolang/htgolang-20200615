package ens

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xxdu521/cmdbgo/agent/config"
	"github.com/imroc/req"
)

type ENS struct {
	conf *config.Config

}

func NewENS(conf *config.Config) *ENS{
	return &ENS{conf:conf}
}

func (s *ENS) Start(){
	logrus.Info("ENS 启动并运行")

	go func(){
		endpoint := fmt.Sprintf("%s/heartbeat/%s/", s.conf.Endpoint, s.conf.UUID)
		for evt := range s.conf.Heartbeat {
			//logrus.Debug("心跳信息")
			//req.Debug = true   //debug模式显示http请求的详细信息
			response, err := req.New().Post(endpoint, req.BodyJSON(evt))
			if err == nil {
				result := map[string]interface{}{}
				response.ToJSON(&result)

				logrus.WithFields(logrus.Fields{
					"result": result,
				}).Debug("心跳发送成功")
			} else {
				logrus.WithFields(logrus.Fields{
					"error": err,
				}).Error("心跳发送失败")
			}
		}
	}()


	go func(){
		endpoint := fmt.Sprintf("%s/register/%s/", s.conf.Endpoint, s.conf.UUID)
		for evt := range s.conf.Register {
			//logrus.Debug("注册信息")
			//req.Debug = true   //debug模式显示http请求的详细信息
			response, err := req.New().Post(endpoint, req.BodyJSON(evt))
			if err == nil {
				result := map[string]interface{}{}
				response.ToJSON(&result)

				logrus.WithFields(logrus.Fields{
					"result": result,
				}).Debug("注册信息成功")
			} else {
				logrus.WithFields(logrus.Fields{
					"error": err,
				}).Error("注册信息失败")
			}
		}
	}()


	go func(){
		endpoint := fmt.Sprintf("%s/log/%s/", s.conf.Endpoint, s.conf.UUID)
		for evt := range s.conf.Log {
			//logrus.Debug("上传日志")
			//req.Debug = true   //debug模式显示http请求的详细信息
			response, err := req.New().Post(endpoint, req.BodyJSON(evt))
			if err == nil {
				result := map[string]interface{}{}
				response.ToJSON(&result)

				logrus.WithFields(logrus.Fields{
					"result": result,
				}).Debug("上传日志成功")
			} else {
				logrus.WithFields(logrus.Fields{
					"error": err,
				}).Error("上传日志失败")
			}
		}
	}()
}
