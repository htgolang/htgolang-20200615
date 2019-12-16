package entity

import (
	"encoding/json"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"time"
)

const (
	LOGResource = 0x0001
)

type Resource struct {
	Load string `json:"load"`
	CPUPrecent float64 `json:"cpu_precent"`
	RAMPrecent float64 `json:"ram_precent"`
	DiskPrecent string `json:"disk_precent"`
}

func NewResource () Resource {
	loadAvgStat,_ := load.Avg()
	load,_ := json.Marshal(loadAvgStat)

	cpuPercents,_ := cpu.Percent(time.Second * 1, false)

	memInfo,_ := mem.VirtualMemory()

	disks := map[string]float64{}
	partitionInfoList,_ := disk.Partitions(true)
	for _,partitionInfo := range partitionInfoList {
		usageInfo,err := disk.Usage(partitionInfo.Device)
		if err != nil {
			continue
		}
		disks[usageInfo.Path] = usageInfo.UsedPercent
	}
	disk,_ := json.Marshal(disks)

	return Resource{
		Load:        string(load),
		CPUPrecent:  cpuPercents[0],
		RAMPrecent:  memInfo.UsedPercent,
		DiskPrecent: string(disk),
	}
}