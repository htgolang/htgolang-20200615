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

type Register struct {
	UUID     string    `json:"uuid"`
	Hostname string    `json:"hostname"`
	IP       string    `json:"ip"`
	OS       string    `json:"os"`
	Arch     string    `json:"arch"`
	CPU      int       `json:"cpu"`
	RAM      int64     `json:"ram"` // MB
	Disk     string    `json:"disk"`
	BootTime time.Time `json:"boottime"`
	Time     time.Time `json:"time"`
}

func NewRegister(uuid string) Register {
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

	partitionInfoList, _ := disk.Partitions(true)

	disks := map[string]int64{}
	for _, partitionInfo := range partitionInfoList {
		usageInfo, err := disk.Usage(partitionInfo.Device)
		if err != nil {
			continue
		}
		disks[usageInfo.Path] = int64(usageInfo.Total / 1024 / 1024 / 1024)
	}

	disk, _ := json.Marshal(disks)
	return Register{
		UUID:     uuid,
		Hostname: hostInfo.Hostname,
		OS:       hostInfo.OS,
		IP:       string(ip),
		Arch:     "", //hostInfo.KernelArch,
		CPU:      cores,
		RAM:      int64(memInfo.Total / 1024 / 1024),
		Disk:     string(disk),
		BootTime: time.Unix(int64(hostInfo.BootTime), 0),
		Time:     time.Now(),
	}
}
