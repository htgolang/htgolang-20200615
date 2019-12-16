package init

import (
	"github.com/xxdu521/cmdbgo/agent/plugins"
	"github.com/xxdu521/cmdbgo/agent/plugins/cycle"
)

func init(){
	plugins.DefaultManager.RegisterCycle(&cycle.Heartbeat{})
	plugins.DefaultManager.RegisterCycle(&cycle.Register{})
	plugins.DefaultManager.RegisterCycle(&cycle.Resource{})
}
