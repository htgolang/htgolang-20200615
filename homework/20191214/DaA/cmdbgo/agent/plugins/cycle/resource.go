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
	p.nextTime = p.nextTime.Add(p.interval)
	return entity.NewLog(p.conf.UUID, entity.LOGResource, entity.NewResource()),nil
}

func (p *Resource) Pipline() chan interface{} {
	return p.conf.Log
}