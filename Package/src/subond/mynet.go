package subond

import (
	"fmt"
	"log"
	"net"
)

func NetPrint() {
	fmt.Println(net.IPv4(8, 8, 8, 8))
}

func NetCIDR() {
	ipv4Addr, ipv4Net, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv4Addr)
	fmt.Println(ipv4Net)

	ipv6Addr, ipv6Net, err := net.ParseCIDR("2001:db8:a0b:12f0::1/32")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv6Addr)
	fmt.Println(ipv6Net)
}

func NetParseIP(ips string) bool {

	ip := net.ParseIP(ips)
	fmt.Println(ip.DefaultMask())
	fmt.Println(ip.DefaultMask().Size())
	fmt.Println(ip.IsGlobalUnicast())
	fmt.Println(net.IPv4Mask(255, 255, 255, 0).Size())
	return true
}

func NetIPNet() {
	ip := net.ParseIP("192.168.12.12")
	mask := ip.DefaultMask()
	var ipn = net.IPNet{ip, mask}
	if ipn.Contains(net.ParseIP("192.168.12.13")) {
		fmt.Println("ok")
	}
	fmt.Println(ipn.Network())
	fmt.Println(ipn.String())
}
