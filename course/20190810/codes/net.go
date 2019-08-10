package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.JoinHostPort("0.0.0.0", "8888"))

	fmt.Println(net.SplitHostPort("127.0.0.1:9999"))

	fmt.Println(net.LookupAddr("61.135.169.121"))
	fmt.Println(net.LookupHost("www.baidu.com"))

	fmt.Println(net.ParseCIDR("192.168.1.1/24"))

	ip := net.ParseIP("::1")
	fmt.Println(ip)
	ips, err := net.LookupIP("www.baidu.com")
	fmt.Println(ips, err)

	ip, ipnet, err := net.ParseCIDR("192.168.1.1/24")

	fmt.Println(ipnet.Contains(net.ParseIP("192.168.1.40")))
	fmt.Println(ipnet.Contains(net.ParseIP("193.168.1.40")))
	fmt.Println(ipnet.Network())

	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		fmt.Println(addr.String(), addr.Network())
	}

	inters, _ := net.Interfaces()

	for _, inter := range inters {
		fmt.Println(inter.Index, inter.Name, inter.MTU, inter.HardwareAddr, inter.Flags)
		fmt.Println(inter.Addrs())
		fmt.Println(inter.MulticastAddrs())
	}

}
