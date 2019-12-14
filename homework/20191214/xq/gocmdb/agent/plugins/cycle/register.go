package cycle

import (
	"github.com/xlotz/gocmdb/agent/config"
	"github.com/xlotz/gocmdb/agent/entity"
	"time"
)

type Register struct {

	conf *config.Config
	interval time.Duration
	nextTime time.Time

}

func (p *Register) Name() string {
	return "register"
}

func (p *Register) Init(config *config.Config)  {
	p.conf = config
	//p.interval = 1 * time.Hour
	p.interval = 10 * time.Second
	p.nextTime = time.Now()

}

func (p *Register) NextTime() time.Time {
	return p.nextTime
}

func (p *Register) Call() (interface{}, error) {

	p.nextTime = p.nextTime.Add(p.interval)

	return entity.NewRegister(p.conf.UUID), nil

}

func (p *Register) Pipline() chan interface{}{
	return p.conf.Register
}