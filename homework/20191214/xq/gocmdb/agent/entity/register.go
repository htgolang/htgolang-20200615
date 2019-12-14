package entity

import (
	"encoding/json"
	"fmt"

	"strings"
	"time"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"

)

type Register struct {
	UUID string `json:"uuid"`
	Hostname string `json:"hostname"`
	IP string `json:"ip"`
	OS string `json:"os"`
	Arch string `json:"arch"`
	CPU int `json:"cpu"`
	MEM int64 `json:"mem"` // MB
	Disk string `json:"disk"`
	BootTime time.Time `json:"boot_time"`
	Time time.Time `json:"time"`
}


func NewRegister(uuid string) Register{
	hostInfo, _ := host.Info()

	ips := []string{}

	inters, _ := net.Interfaces()

	for _, inter := range inters {
		for _, addr := range inter.Addrs{

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

	cpus, _ := cpu.Info()

	coses:=0
	for _, cpuInfo:= range cpus {
		coses += int(cpuInfo.Cores)
	}
	memInfo,_ := mem.VirtualMemory()

	diskinfolist,_:= disk.Partitions(true)

	disks := map[string]int64{}

	for _, diskinfo := range diskinfolist{
		usedisk, err := disk.Usage(diskinfo.Device)
		if err != nil {
			continue
		}
		//fmt.Println(usedisk, err)
		disks[usedisk.Path] = int64(usedisk.Total/1024/1024/1024)
		fmt.Println(usedisk.Total/1024/1024/1024)
	}

	disk,_ := json.Marshal(disks)

	return Register{
		UUID:	uuid,
		Hostname:	hostInfo.Hostname,
		OS:		hostInfo.OS,
		Arch:	hostInfo.KernelArch,
		IP: 	string(ip),
		CPU:	coses,
		MEM:	int64(memInfo.Total/1024/1024),
		Disk:	string(disk),
		BootTime:	time.Unix(int64(hostInfo.BootTime), 0),
		Time:	time.Now(),


	}
}