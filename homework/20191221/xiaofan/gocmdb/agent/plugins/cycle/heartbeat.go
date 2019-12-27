package cycle

import (
	"github.com/dcosapp/gocmdb/agent/config"
	"github.com/dcosapp/gocmdb/agent/entity"
	"time"
)

type Heartbeat struct {
	conf     *config.Config
	interval time.Duration
	nextTime time.Time
}

// 获取插件名称
func (p *Heartbeat) Name() string {
	return "heartbeat"
}

// 将config里面的配置传给插件
func (p *Heartbeat) Init(conf *config.Config) {
	p.conf = conf
	p.interval = 10 * time.Second
	p.nextTime = time.Now()
}

// 返回当前时间（程序启动立即执行一次）
func (p *Heartbeat) NextTime() time.Time {
	return p.nextTime
}

// return Heartbeat{ UUID: p.conf.UUID, Time: time.Now(), }
func (p *Heartbeat) Call() (interface{}, error) {
	p.nextTime = p.nextTime.Add(p.interval)
	return entity.NewHeartbeat(p.conf.UUID), nil
}

// 定义该插件使用哪个chan进行数据传递
func (p *Heartbeat) Pipline() chan interface{} {
	return p.conf.Heartbeat
}
