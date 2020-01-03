package task

import (
	"github.com/imsilence/gocmdb/agent/config"
	"github.com/shirou/gopsutil/process"
)

type Process struct {
	conf *config.Config
}

func (p *Process) Name() string {
	return "process"
}

func (p *Process) Init(conf *config.Config) {
	p.conf = conf
}

func (p *Process) Call(params string) (interface{}, error) {
	pids, err := process.Pids()
	if err != nil {
		return nil, err
	}
	rs := make([]map[string]interface{}, len(pids))
	for index, pid := range pids {
		ps, err := process.NewProcess(pid)
		if err != nil {
			continue
		}
		name, _ := ps.Name()
		exe, _ := ps.Exe()
		cmd, _ := ps.Cmdline()
		createdTime, _ := ps.CreateTime()
		ppid, _ := ps.Ppid()
		cwd, _ := ps.Cwd()
		numFDs, _ := ps.NumFDs()

		numThreads, _ := ps.NumThreads()
		memoryInfo, _ := ps.MemoryInfo()
		connections, _ := ps.Connections()
		rs[index] = map[string]interface{}{
			"pid":         pid,
			"ppid":        ppid,
			"name":        name,
			"exe":         exe,
			"cmd":         cmd,
			"createdTime": createdTime,
			"cwd":         cwd,
			"numFDs":      numFDs,
			"numThreads":  numThreads,
			"memoryInfo":  memoryInfo,
			"connections": connections,
		}
	}
	return rs, nil
}
