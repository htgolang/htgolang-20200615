package init

import (
	"github.com/imsilence/gocmdb/agent/plugins"
	"github.com/imsilence/gocmdb/agent/plugins/cycle"
)

func init() {
	plugins.DefaultManager.RegisterCycle(&cycle.Heartbeat{})
	plugins.DefaultManager.RegisterCycle(&cycle.Register{})
	plugins.DefaultManager.RegisterCycle(&cycle.Resource{})
}
