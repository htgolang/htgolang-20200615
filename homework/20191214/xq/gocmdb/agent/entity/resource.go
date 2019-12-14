package entity

import (
	"encoding/json"

	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
	"time"
)

type Resource struct {
	CPUPrecent float64 `json:"cpu_precent"`
	MEMPrecent float64 `json:"mem_precent"`
	DiskPrecent string `json:"disk_precent"`
	Load string `json:"load"`
}

func NewResource() Resource {

	loadAvgStat , _ := load.Avg()

	load,_ := json.Marshal(loadAvgStat)

	cpuPercents, _ :=cpu.Percent(time.Second, false)

	memPercents, _ := mem.VirtualMemory()

	diskinfolist,_:= disk.Partitions(true)

	disks := map[string]float64{}

	for _, diskinfo := range diskinfolist{
		usedisk, err := disk.Usage(diskinfo.Device)
		if err != nil {
			continue
		}
		//fmt.Println(usedisk, err)
		disks[usedisk.Path] = float64(usedisk.Total/1024/1024/1024)

	}

	disk,_ := json.Marshal(disks)

	return Resource{
		Load: string(load),
		CPUPrecent:cpuPercents[0],
		MEMPrecent:memPercents.UsedPercent,
		DiskPrecent:string(disk),

	}
}
