package cycle

import (
	"github.com/dcosapp/gocmdb/agent/config"
	"github.com/dcosapp/gocmdb/agent/entity"
	"time"
)

type Register struct {
	conf     *config.Config
	interval time.Duration
	nextTime time.Time
}

// 获取插件名称
func (p *Register) Name() string {
	return "register"
}

// 将config里面的配置传给插件
func (p *Register) Init(conf *config.Config) {
	p.conf = conf
	p.interval = time.Hour
	p.nextTime = time.Now()
}

// 返回当前时间（程序启动立即执行一次）
func (p *Register) NextTime() time.Time {
	return p.nextTime
}

// return Register{...}
func (p *Register) Call() (interface{}, error) {
	p.nextTime = p.nextTime.Add(p.interval)
	return entity.NewRegister(p.conf.UUID), nil
}

// 定义该插件使用哪个chan进行数据传递
func (p *Register) Pipline() chan interface{} {
	return p.conf.Register
}
