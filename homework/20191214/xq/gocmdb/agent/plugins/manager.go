package plugins

import (

	//"database/sql/driver"
	"github.com/xlotz/gocmdb/agent/config"
	"github.com/sirupsen/logrus"
	"time"
)

type Manager struct {
	Cycles map[string]CyclePlugin
}

func NewManager() *Manager {
	return &Manager{
		Cycles:make(map[string]CyclePlugin),
	}
}

//注册
func (m *Manager) RegisterCycle(p CyclePlugin){
	m.Cycles[p.Name()] = p
	logrus.WithFields(logrus.Fields{
		"Name": p.Name(),

	}).Info("插件注册成功")
}

// config new 完后，调用

func (m *Manager) Init(config *config.Config){
	for key, plugin := range m.Cycles {
		plugin.Init(config)

		logrus.WithFields(logrus.Fields{
			"Name": key,
			}).Info("初始化插件成功")
	}
}

func (m *Manager) Start(){
	// 启动历程
	go m.StartCycle()

}

func (m *Manager) StartCycle(){
	for now := range time.Tick(time.Second){
		for name, plugin := range m.Cycles{

			//now > plugin.NextTime()
			if now.After(plugin.NextTime()) {

				if result, err := plugin.Call(); err == nil {

					logrus.WithFields(logrus.Fields{
						"Name": name,
						"Result": result,
					}).Debug("插件执行")

					plugin.Pipline() <- result
				}else {

					logrus.WithFields(logrus.Fields{
						"Name": name,
						"Error": err,
					}).Debug("插件执行失败")

				}


			}
		}
	}
}

var  DefaultManager  = NewManager()