package plugins

import (
	"errors"
	"time"

	"github.com/imsilence/gocmdb/agent/config"
	"github.com/imsilence/gocmdb/agent/entity"
	"github.com/sirupsen/logrus"
)

type Manager struct {
	Conf   *config.Config
	Cycles map[string]CyclePlugin
	Tasks  map[string]TaskPlugin
}

func NewManager() *Manager {
	return &Manager{
		Cycles: make(map[string]CyclePlugin),
		Tasks:  make(map[string]TaskPlugin),
	}
}

func (m *Manager) RegisterCycle(p CyclePlugin) {
	m.Cycles[p.Name()] = p
	logrus.WithFields(logrus.Fields{
		"Name": p.Name(),
		"Type": "周期型",
	}).Info("插件注册")
}

func (m *Manager) RegisterTask(p TaskPlugin) {
	m.Tasks[p.Name()] = p
	logrus.WithFields(logrus.Fields{
		"Name": p.Name(),
		"Type": "任务型",
	}).Info("插件注册")
}

func (m *Manager) Init(conf *config.Config) {
	m.Conf = conf
	for name, plugin := range m.Cycles {
		plugin.Init(conf)
		logrus.WithFields(logrus.Fields{
			"Name": name,
		}).Info("初始化插件")
	}
	for name, plugin := range m.Tasks {
		plugin.Init(conf)
		logrus.WithFields(logrus.Fields{
			"Name": name,
		}).Info("初始化插件")
	}
}

func (m *Manager) Start() {
	go m.StartCycle()
	go m.StartTask()
}

func (m *Manager) StartCycle() {
	for now := range time.Tick(time.Second) {
		for name, plugin := range m.Cycles {
			if now.After(plugin.NextTime()) {
				go func(pluginName string, plugin CyclePlugin) {
					if evt, err := plugin.Call(); err == nil {
						logrus.WithFields(logrus.Fields{
							"Name":   pluginName,
							"Result": evt,
						}).Debug("插件执行")
						plugin.Pipline() <- evt
					} else {
						logrus.WithFields(logrus.Fields{
							"Name":  pluginName,
							"error": err,
						}).Debug("插件执行失败")
					}
				}(name, plugin)

			}
		}
	}
}

func (m *Manager) StartTask() {
	for task := range m.Conf.Task {
		taskObj, _ := task.(entity.Task)
		if plugin, ok := m.Tasks[taskObj.Plugin]; !ok {
			logrus.WithFields(logrus.Fields{
				"task": taskObj,
			}).Error("插件执行失败, 插件不存在")
			m.Conf.TaskResult <- entity.NewResult(taskObj, nil, errors.New("插件不存在"))
		} else {
			go func(pluginName string, plugin TaskPlugin) {
				result, err := plugin.Call(taskObj.Params)
				logrus.WithFields(logrus.Fields{
					"Name":   pluginName,
					"task":   taskObj,
					"Result": result,
					"Err":    err,
				}).Error("插件执行完成")

				m.Conf.TaskResult <- entity.NewResult(taskObj, result, err)
			}(plugin.Name(), plugin)

		}
	}
}

var DefaultManager = NewManager()
