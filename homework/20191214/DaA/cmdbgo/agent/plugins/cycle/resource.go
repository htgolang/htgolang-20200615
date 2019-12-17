package cycle

import (
	"github.com/xxdu521/cmdbgo/agent/config"
	"github.com/xxdu521/cmdbgo/agent/entity"
	"time"
)

type Resource struct {
	conf *config.Config
	interval time.Duration
	nextTime time.Time
}
func (p *Resource) Name() string {
	return "log"
}
func (p *Resource) Init(conf *config.Config) {
	p.conf = conf
	p.interval = 1 * time.Minute
	p.nextTime = time.Now()
}
func (p *Resource) NextTime() time.Time {
	return p.nextTime
}
func (p *Resource) Call() (interface{}, error){
	p.nextTime = p.nextTime.Add(p.interval)  //time.Time.Add方法，把一个时间加上一个时间范围(time.Duration),例如 2019-01-02 03:04:05 + 60秒
	return entity.NewLog(p.conf.UUID, entity.LOGResource, entity.NewResource()),nil
} //resource插件，使用log通道的msg字段，上报信息。
func (p *Resource) Pipline() chan interface{} {
	return p.conf.Log
} //返回一个通道，指定resource的插件，走Log通道上报数据