package plugins

import (
	"github.com/dcosapp/gocmdb/agent/config"
	"github.com/sirupsen/logrus"
	"time"
)

// 插件管理
type Manager struct {
	Cycles map[string]CyclePlugin
}

func NewManager() *Manager {
	return &Manager{
		Cycles: make(map[string]CyclePlugin),
	}
}

// 注册插件，即将插件放入插件管理的结构体中
func (m *Manager) RegisterCycle(p CyclePlugin) {
	m.Cycles[p.Name()] = p
	logrus.WithFields(logrus.Fields{"Name": p.Name(),}).Info("插件管理器: 插件注册")
}

// 初始化插件
func (m *Manager) Init(conf *config.Config) {
	for name, plugin := range m.Cycles {
		// 调用插件自己的初始化
		plugin.Init(conf)
		logrus.WithFields(logrus.Fields{"Name": name,}).Info("插件管理器: 初始化插件")
	}
}

// 调用startCycle启动插件
func (m *Manager) Start() {
	go m.startCycle()
}

func (m *Manager) startCycle() {
	// 每秒钟执行一次
	for now := range time.Tick(time.Second) {
		// 遍历插件
		for name, plugin := range m.Cycles {
			// 判断当前时间是否大于插件的执行时间
			if now.After(plugin.NextTime()) {
				// 大于则执行插件
				if evt, err := plugin.Call(); err == nil {
					logrus.WithFields(logrus.Fields{"Name": name, "Result": evt,}).Debug("插件执行")
					// 将产生的消息送给相对应的chan,由ens接收并处理消息
					plugin.Pipline() <- evt
				} else {
					logrus.WithFields(logrus.Fields{"Name": name, "Error": err.Error(),}).Error("插件执行失败")
				}
			}
		}
	}
}

var DefaultManager = NewManager()
