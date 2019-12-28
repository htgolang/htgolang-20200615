package plugins

import (
	"time"
	"github.com/sirupsen/logrus"
	"github.com/imsilence/gocmdb/agent/config"
)

type Manager struct {
	Cycles map[string]CyclePlugin
}

func NewManager() *Manager {
	return &Manager{
		Cycles: make(map[string]CyclePlugin),
	}
}

func (m *Manager) RegisterCycle(p CyclePlugin) {
	m.Cycles[p.Name()] = p
	logrus.WithFields(logrus.Fields{
		"Name" : p.Name(),
	}).Info("插件注册")
}

func (m *Manager) Init(conf *config.Config) {
	for name, plugin := range m.Cycles {
		plugin.Init(conf)
		logrus.WithFields(logrus.Fields{
			"Name": name,
		}).Info("初始化插件")
	}
}

func (m *Manager) Start() {
	go m.StartCycle()
}

func (m *Manager) StartCycle() {
	for now := range time.Tick(time.Second) {
		for name, plugin := range m.Cycles {
			if now.After(plugin.NextTime()) {
				if evt, err := plugin.Call(); err == nil {
					logrus.WithFields(logrus.Fields{
						"Name" : name,
						"Result" : evt,
					}).Debug("插件执行")
					plugin.Pipline() <- evt
				} else {
					logrus.WithFields(logrus.Fields{
						"Name" : name,
						"error" : err,
					}).Debug("插件执行失败")
				}
			}
		}
	}
}

var DefaultManager = NewManager()