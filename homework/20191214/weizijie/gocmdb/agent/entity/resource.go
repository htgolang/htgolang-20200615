package entity

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

type Resource struct {
	Load        string  `json:"load"`
	CPUPrecent  float64 `json:"cpu_precent"`
	RAMPrecent  float64 `json:"ram_precent"`
	DiskPrecent string  `json:"disk_precent"`
}

func NewResource() Resource {
	loadAvgStat, _ := load.Avg()
	cpuPercents, _ := cpu.Percent(time.Second, false)
	memInfo, _ := mem.VirtualMemory()

	disks := map[string]float64{}
	partitionInfoList, _ := disk.Partitions(true)
	for _, partitionInfo := range partitionInfoList {
		usageInfo, _ := disk.Usage(partitionInfo.Device)
		disks[usageInfo.Path] = usageInfo.UsedPercent
	}

	disk, _ := json.Marshal(disks)
	fmt.Println(string(disk))
	load, _ := json.Marshal(loadAvgStat)
	return Resource{
		Load:        string(load),
		CPUPrecent:  cpuPercents[0],
		RAMPrecent:  memInfo.UsedPercent,
		DiskPrecent: string(disk),
	}
}
