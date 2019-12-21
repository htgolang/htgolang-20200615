package entity

import (
	"encoding/json"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

type Resource struct {
	Load        string  `json:"load"`
	CPUPrecent  float64 `json:"cpu_percent"`
	RAMPrecent  float64 `json:"ram_percent"`
	DiskPrecent string  `json:"disk_percent"`
}

func NewResource() Resource {
	loadAvgStat, _ := load.Avg()
	cpuPercents, _ := cpu.Percent(time.Second, false)

	memInfo, _ := mem.VirtualMemory()

	disks := map[string]float64{}

	partitionInfoList, _ := disk.Partitions(true)
	for _, partitionInfo := range partitionInfoList {
		usageInfo, err := disk.Usage(partitionInfo.Device)
		if err != nil {
			continue
		}
		disks[usageInfo.Path] = usageInfo.UsedPercent
	}

	load, _ := json.Marshal(loadAvgStat)
	disk, _ := json.Marshal(disks)
	return Resource{
		Load:       string(load),
		CPUPrecent: cpuPercents[0],
		RAMPrecent: memInfo.UsedPercent,
		DiskPrecent: string(disk),
	}
}
