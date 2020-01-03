package ens

import (
	"fmt"
	"time"

	"github.com/imroc/req"
	"github.com/sirupsen/logrus"

	"github.com/imsilence/gocmdb/agent/config"
	"github.com/imsilence/gocmdb/agent/entity"
)

type ENS struct {
	conf *config.Config
}

func NewENS(conf *config.Config) *ENS {
	return &ENS{conf: conf}
}

func (s *ENS) Start() {
	logrus.Info("ENS 开始运行")

	headers := req.Header{"Token": s.conf.Token}
	request := req.New()

	go func() {
		endpoint := fmt.Sprintf("%s/heartbeat/%s/", s.conf.Endpoint, s.conf.UUID)
		for evt := range s.conf.Heartbeat {
			response, err := request.Post(endpoint, req.BodyJSON(evt), headers)
			if err == nil {
				result := map[string]interface{}{}
				response.ToJSON(&result)
				logrus.WithFields(logrus.Fields{
					"heartbeat": evt,
					"result":    result,
				}).Debug("上传心跳信息成功")
			} else {
				logrus.WithFields(logrus.Fields{
					"heartbeat": evt,
					"error":     err,
				}).Error("上传心跳信息失败")
			}
		}
	}()

	go func() {
		endpoint := fmt.Sprintf("%s/register/%s/", s.conf.Endpoint, s.conf.UUID)
		for evt := range s.conf.Register {
			response, err := request.Post(endpoint, req.BodyJSON(evt), headers)
			if err == nil {
				result := map[string]interface{}{}
				response.ToJSON(&result)
				logrus.WithFields(logrus.Fields{
					"info":   evt,
					"result": result,
				}).Debug("注册成功")
			} else {
				logrus.WithFields(logrus.Fields{
					"info":  evt,
					"error": err,
				}).Error("注册失败")
			}
		}
	}()

	go func() {
		endpoint := fmt.Sprintf("%s/log/%s/", s.conf.Endpoint, s.conf.UUID)
		for evt := range s.conf.Log {
			response, err := request.Post(endpoint, req.BodyJSON(evt), headers)
			if err == nil {
				result := map[string]interface{}{}
				response.ToJSON(&result)
				logrus.WithFields(logrus.Fields{
					"log":    evt,
					"result": result,
				}).Debug("日志上传成功")
			} else {
				logrus.WithFields(logrus.Fields{
					"log":   evt,
					"error": err,
				}).Error("日志上传失败")
			}
		}
	}()

	go func() {
		endpoint := fmt.Sprintf("%s/result/%s/", s.conf.Endpoint, s.conf.UUID)
		for evt := range s.conf.TaskResult {
			response, err := request.Post(endpoint, req.BodyJSON(evt), headers)
			if err == nil {
				result := map[string]interface{}{}
				response.ToJSON(&result)
				logrus.WithFields(logrus.Fields{
					"taskResult": evt,
					"result":     result,
				}).Debug("任务结果上传成功")
			} else {
				logrus.WithFields(logrus.Fields{
					"taskResult": evt,
					"error":      err,
				}).Error("任务结果上传失败")
			}
		}
	}()

	go func() {
		endpoint := fmt.Sprintf("%s/task/%s/", s.conf.Endpoint, s.conf.UUID)
		for now := range time.Tick(10 * time.Second) {
			response, err := request.Get(endpoint, req.QueryParam{"time": now.Unix()}, headers)
			if err == nil {
				result := map[string]interface{}{}
				response.ToJSON(&result)
				logrus.WithFields(logrus.Fields{
					"result": result,
				}).Debug("获取任务成功")
				tasks, _ := result["result"].([]interface{})
				for _, taskMap := range tasks {
					if task, err := entity.NewTask(taskMap); err == nil {
						s.conf.Task <- task
					}
				}
			} else {
				logrus.WithFields(logrus.Fields{
					"error": err,
				}).Error("获取任务失败")
			}
		}
	}()
}
