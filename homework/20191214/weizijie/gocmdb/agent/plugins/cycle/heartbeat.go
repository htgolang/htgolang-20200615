package cycle

import (
	"time"

	"github.com/JevonWei/gocmdb/agent/config"
	"github.com/JevonWei/gocmdb/agent/entity"
)

type Heartbeat struct {
	conf     *config.Config
	interval time.Duration
	nextTime time.Time
}

func (p *Heartbeat) Name() string {
	return "heartbeat"
}

func (p *Heartbeat) Init(conf *config.Config) {
	p.conf = conf
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

func (p *Heartbeat) Pipline() chan interface{} {
	return p.conf.Heartbeat
}
