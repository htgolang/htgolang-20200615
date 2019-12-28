package init

import (
	"github.com/xlotz/gocmdb/agent/plugins"
	"github.com/xlotz/gocmdb/agent/plugins/cycle"
)

func init()  {
	plugins.DefaultManager.RegisterCycle(&cycle.Heartbeat{})
	plugins.DefaultManager.RegisterCycle(&cycle.Register{})
	plugins.DefaultManager.RegisterCycle(&cycle.Resource{})
}
