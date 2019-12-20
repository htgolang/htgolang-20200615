package init

import (
	"github.com/JevonWei/gocmdb/agent/plugins"
	"github.com/JevonWei/gocmdb/agent/plugins/cycle"
)

func init() {
	plugins.DefaultManager.RegisterCycle(&cycle.Heartbeat{})
	plugins.DefaultManager.RegisterCycle(&cycle.Register{})
	plugins.DefaultManager.RegisterCycle(&cycle.Resource{})
}
