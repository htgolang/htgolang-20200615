package entity

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type Registry struct {
	UUID     string    `json:"uuid"`
	Hostname string    `json:"hostname"`
	IP       string    `json:"ip"`
	OS       string    `json:"os"`
	Arch     string    `json:"arch"`
	CPU      int       `json:"cpu"`
	RAM      int64     `json:"ram"` //MB
	Disk     string    `json:"disk"`
	BootTime time.Time `json:"boottime"`
	Time     time.Time `json:"time"`
}

func NewRegister(uuid string) Registry {
	hostInfo, _ := host.Info()
	ips := []string{}
	interfaceStatList, _ := net.Interfaces()
	for _, interfaceStat := range interfaceStatList {
		for _, addr := range interfaceStat.Addrs {
			if strings.Index(addr.Addr, ":") >= 0 {
				continue
			}
			if strings.Index(addr.Addr, "127.") == 0 {
				continue
			}
			nodes := strings.Split(addr.Addr, "/")
			ips = append(ips, nodes[0])
		}
	}
	ip, _ := json.Marshal(ips)

	cores := 0
	cpuInfoList, _ := cpu.Info()
	for _, cpuInfo := range cpuInfoList {
		cores += int(cpuInfo.Cores)
	}

	memInfo, _ := mem.VirtualMemory()

	disks := map[string]int64{}
	partitionInfoList, _ := disk.Partitions(true)
	for _, partitionInfo := range partitionInfoList {
		usageInfo, _ := disk.Usage(partitionInfo.Device)
		disks[usageInfo.Path] = int64(usageInfo.Total / 1024 / 1024 / 1024) // GB
	}
	disk, _ := json.Marshal(disks)

	return Registry{
		UUID:     uuid,
		Hostname: hostInfo.Hostname,
		OS:       hostInfo.OS,
		IP:       string(ip),
		Arch:     "", //hostInfo.KernelArch,
		CPU:      cores,
		RAM:      int64(memInfo.Total / 1024 / 1024), // MB
		Disk:     string(disk),
		BootTime: time.Unix(int64(hostInfo.BootTime), 0),
		Time:     time.Now(),
	}
}
