package plugins

import (
	"time"
	"github.com/imsilence/gocmdb/agent/config"
)

type CyclePlugin interface {
	Name() string
	Init(*config.Config)
	NextTime() time.Time
	Call() (interface{}, error)
	Pipline() chan interface{}
}