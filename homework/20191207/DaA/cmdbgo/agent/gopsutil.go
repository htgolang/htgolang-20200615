package main

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/disk"
	"time"
)

func main(){
	hostinfo,_ := host.Info()
	fmt.Println(hostinfo)

	cpuinfos,_ := cpu.Info()
	for _,cpuinfo := range cpuinfos {
		fmt.Println(cpuinfo)
	}

	cpupercents,_  := cpu.Percent(time.Second, false)
	fmt.Println(cpupercents)
	cpupercents,_  = cpu.Percent(time.Second, true)
	fmt.Println(cpupercents)

	meminfo,_ := mem.VirtualMemory()
	fmt.Println(meminfo)

	interfaces,_ := net.Interfaces()
	for _,interinfo := range interfaces {
		fmt.Println(interinfo)
	}

	fmt.Println(net.IOCounters(false))
	fmt.Println(net.IOCounters(true))

	loadinfo,_ := load.Avg()
	fmt.Println(loadinfo)


	diskinfos,_ := disk.Partitions(true)
	for _,diskinfo := range diskinfos {
		fmt.Println(diskinfo)
	}

	diskinfos,_ = disk.Partitions(false)
	for _,diskinfo := range diskinfos {
		fmt.Println(diskinfo)
		fmt.Println(disk.Usage(diskinfo.Device))
	}

	diskcounters,_ := disk.IOCounters()
	fmt.Println(diskcounters)

}
