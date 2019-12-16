package cycle

import (
	"github.com/xlotz/gocmdb/agent/config"
	"github.com/xlotz/gocmdb/agent/entity"
	"time"
)

type Heartbeat struct {

	conf *config.Config
	interval time.Duration
	nextTime time.Time

}

func (p *Heartbeat) Name() string {
	return "heartbeat"
}

func (p *Heartbeat) Init(config *config.Config)  {
	p.conf = config
	p.interval = 10 * time.Second

	p.nextTime = time.Now()

}

func (p *Heartbeat) NextTime() time.Time {
	return p.nextTime
}

func (p *Heartbeat) Call() (interface{}, error) {

	p.nextTime = p.nextTime.Add(p.interval)

	return entity.NewHeartbeat(p.conf.UUID), nil

}

func (p *Heartbeat) Pipline() chan interface{}{
	return p.conf.Heartbeat
}