package cycle

import (
	"github.com/dcosapp/gocmdb/agent/config"
	"github.com/dcosapp/gocmdb/agent/entity"
	"time"
)

type Resource struct {
	conf     *config.Config
	interval time.Duration
	nextTime time.Time
}

// 获取插件名称
func (p *Resource) Name() string {
	return "resource"
}

// 获取插件名称
func (p *Resource) Init(conf *config.Config) {
	p.conf = conf
	//p.interval = time.Hour
	p.interval = time.Minute
	p.nextTime = time.Now()
}

// 返回当前时间（程序启动立即执行一次）
func (p *Resource) NextTime() time.Time {
	return p.nextTime
}

// return return Log{ UUID: p.conf.UUID, Type: entity.LOGResource, Msg: string(json.Marshal(Resource{...})), Time: time.Now(),}
func (p *Resource) Call() (interface{}, error) {
	p.nextTime = p.nextTime.Add(p.interval)
	return entity.NewLog(p.conf.UUID, entity.LOGResource, entity.NewResource()), nil
}

// 定义该插件使用哪个chan进行数据传递
func (p *Resource) Pipline() chan interface{} {
	return p.conf.Log
}
