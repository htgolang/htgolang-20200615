package main

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
)


func main() {
	pids, _ := process.Pids()
	for _, pid := range pids {
		ps, _ := process.NewProcess(pid)
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
		fmt.Println(map[string]interface{}{
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
		})
	}
}
