package cycle

import (
	"github.com/xlotz/gocmdb/agent/config"
	"github.com/xlotz/gocmdb/agent/entity"
	"time"
)

type Resource struct {

	conf *config.Config
	interval time.Duration
	nextTime time.Time

}

func (p *Resource) Name() string {
	return "resource"
}

func (p *Resource) Init(config *config.Config)  {
	p.conf = config
	p.interval = 1 * time.Minute
	//p.interval = 10 * time.Second
	p.nextTime = time.Now()

}

func (p *Resource) NextTime() time.Time {
	return p.nextTime
}

func (p *Resource) Call() (interface{}, error) {

	p.nextTime = p.nextTime.Add(p.interval)

	return entity.NewLog(p.conf.UUID ,entity.LOGResource, entity.NewResource()), nil

}

func (p *Resource) Pipline() chan interface{}{
	return p.conf.Log
}