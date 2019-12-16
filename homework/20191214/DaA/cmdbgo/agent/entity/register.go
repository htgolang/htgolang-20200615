package entity

import (
	"encoding/json"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/disk"
	"strings"
	"time"
	)

type Register struct {
	UUID 		string 	`json:"uuid"`
	HostName 	string 	`json:"host_name"`
	IP 			string 	`json:"ip"`
	OS 			string 	`json:"os"`
	Arch 		string 	`json:"arch"`
	CPU 		int		`json:"cpu"`
	RAM 		int64 	`json:"ram"` //MB
	Disk 		string 	`json:"disk"`
	BootTime 	time.Time `json:"boot_time"`
	Time 		time.Time `json:"time"`
}

func NewRegister(uuid string) Register{
	hostinfo,_ := host.Info()

	cores := 0
	cpuinfoList,_ := cpu.Info()
	for _,cpuinfo := range cpuinfoList {
		cores += int(cpuinfo.Cores)
	}

	meminfo,_ := mem.VirtualMemory()

	disks := map[string]int64{}
	partitionInfoList,_ := disk.Partitions(true)
	for _,partitionInfo := range partitionInfoList {
		usageInfo,err := disk.Usage(partitionInfo.Device)
		if err == nil {
			disks[usageInfo.Path] = int64(usageInfo.Total / 1024 / 1024 / 1024)
		}
	}
	disk,_ := json.Marshal(disks)

	ips := []string{}
	interfaceStatList,_ := net.Interfaces()
	for _,interfaceStat := range interfaceStatList{
		for _,addr := range interfaceStat.Addrs {
			if strings.Index(addr.Addr,":") >=0 {
				continue
			}
			if strings.Index(addr.Addr,"127.") == 0 {
				continue
			}
			nodes := strings.Split(addr.Addr,"/")
			ips = append(ips, nodes[0])
		}
	}
	ip,_ := json.Marshal(ips)

	return Register{
		UUID: uuid,
		HostName: hostinfo.Hostname,
		IP: string(ip),
		OS: hostinfo.OS,
		Arch: hostinfo.KernelArch,
		CPU: cores,
		RAM: int64(meminfo.Total),
		Disk: string(disk),
		BootTime: time.Unix(int64(hostinfo.BootTime),0),
		Time: time.Now(),
	}
}
