package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func main() {
	hostInfo, _ := host.Info()
	fmt.Printf("%#v\n", hostInfo)

	cpuInfos, _ := cpu.Info()
	for _, cpuInfo := range cpuInfos {
		fmt.Printf("%#v\n", cpuInfo)
	}

	cpuPercents, _ := cpu.Percent(time.Second, false)
	fmt.Println(cpuPercents)
	cpuPercents, _ = cpu.Percent(time.Second, true)
	fmt.Println(cpuPercents)

	memInfo, _ := mem.VirtualMemory()
	fmt.Println(memInfo)

	interInfos, _ := net.Interfaces()
	for _, interInfo := range interInfos {
		fmt.Println(interInfo)
	}

	fmt.Println(net.IOCounters(false))
	fmt.Println(net.IOCounters(true))

	loadInfo, _ := load.Avg()
	fmt.Println(loadInfo)

	diskInfos, _ := disk.Partitions(false)
	for _, diskInfo := range diskInfos {
		fmt.Println(diskInfo)
		diskUsage, _ := disk.Usage(diskInfo.Device)
		fmt.Println(diskUsage)
	}

	diskCounters, _ := disk.IOCounters()
	fmt.Println(diskCounters)
}
