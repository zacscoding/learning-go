package main

import (
	"fmt"
	"net"
)

// 네트워크 정보를 조회한다
func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, i := range interfaces {
		fmt.Printf("Interface:%v\n", i.Name)
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		addresses, err := byName.Addrs()
		for k, v := range addresses {
			fmt.Printf("Interface address #%v: %v\n", k, v.String())
		}
		fmt.Println()
	}
}
