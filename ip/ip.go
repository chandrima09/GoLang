//2. ip(1)
//   Create a program which lists information about network interfaces:
//      * interface name
//      * MAC address
//      * IP addresses
//```
//{Index:1 MTU:65536 Name:lo HardwareAddr: Flags:up|loopback}
//  [127.0.0.1/32]
//```

package main

import (
    "fmt"
    "net"
)

func main() {
	interfaces, _ := net.Interfaces()
	for _, inter := range interfaces {
		fmt.Println("Index:",inter.Index, "  MTU:",inter.MTU, "  Name:",inter.Name, "  HardwareAddr:",inter.HardwareAddr, "  Flags:",inter.Flags)
		addrs,_ := inter.Addrs()
		for _, ipaddr := range addrs {
			fmt.Println("Addr:", ipaddr)
		}
	}
}
