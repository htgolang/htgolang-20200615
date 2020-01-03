package init

import (
	"github.com/imsilence/gocmdb/agent/plugins"
	"github.com/imsilence/gocmdb/agent/plugins/task"
)

func init() {
	plugins.DefaultManager.RegisterTask(&task.Process{})
}
