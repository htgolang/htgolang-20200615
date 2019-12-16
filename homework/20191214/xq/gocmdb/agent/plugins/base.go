package plugins

import (
	"github.com/xlotz/gocmdb/agent/config"

	"time"
)

type CyclePlugin interface {
	Name() string
	Init(*config.Config)
	NextTime() time.Time
	Call() (interface{}, error)
	Pipline() chan interface{}  // 和插件绑定
}
