package entity

import (
	"encoding/json"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"time"
)

type Resource struct {
	Load        string  `json:"load"`
	CPUPercent  float64 `json:"cpu_percent"`
	RAMPercent  float64 `json:"ram_percent"`
	DiskPercent string  `json:"disk_percent"`
}

// 返回资源信息给 entity/log 进行处理
func NewResource() Resource {
	loadAvgStat, _ := load.Avg()
	load, _ := json.Marshal(loadAvgStat)

	cpuPercents, _ := cpu.Percent(time.Second, false)

	memInfo, _ := mem.VirtualMemory()

	partitionInfoList, _ := disk.Partitions(true)

	disks := map[string]float64{}
	for _, partitionInfo := range partitionInfoList {
		usageInfo, err := disk.Usage(partitionInfo.Device)
		if err != nil {
			continue
		}
		disks[usageInfo.Path] = usageInfo.UsedPercent
	}
	disk, _ := json.Marshal(disks)
	return Resource{
		Load:        string(load),
		CPUPercent:  cpuPercents[0],
		RAMPercent:  memInfo.UsedPercent,
		DiskPercent: string(disk),
	}
}
