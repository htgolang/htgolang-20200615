package cycle

import (
	"github.com/xxdu521/cmdbgo/agent/config"
	"github.com/xxdu521/cmdbgo/agent/entity"
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

func (p *Register) Init(conf *config.Config) {
	p.conf = conf
	p.interval = 1 * time.Hour
	//p.interval = 10 * time.Second  //调试，10秒一次
	p.nextTime = time.Now()
}

func (p *Register) NextTime() time.Time {
	return p.nextTime
}

func (p *Register) Call() (interface{}, error){
	p.nextTime = p.nextTime.Add(p.interval)
	return entity.NewRegister(p.conf.UUID),nil
}

func (p *Register) Pipline() chan interface{} {
	return p.conf.Register
}