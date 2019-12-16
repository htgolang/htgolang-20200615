package plugins

import (
	"github.com/xxdu521/cmdbgo/agent/config"
	"time"
)

type CyclePlugin interface {
	Name()string
	Init(*config.Config)
	NextTime() time.Time
	Call() (interface{}, error)
	Pipline() chan interface{}
}
