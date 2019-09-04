package main

import (
	"fmt"
	"net"
	"os"
)

func lookIP(address string) ([]string, error) {
	// /etc/hosts 파일을 조회
	hosts, err := net.LookupAddr(address)
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

func lookHostname(hostname string) ([]string, error) {
	IPs, err := net.LookupHost(hostname)
	if err != nil {
		return nil, err
	}
	return IPs, nil
}

// IP 주소라면 이를 그대로 사용하고, 아니면 IP 주소로 변환
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}

	input := arguments[1]
	IPaddress := net.ParseIP(input)

	if IPaddress == nil {
		IPs, err := lookHostname(input)
		if err == nil {
			for _, singleIP := range IPs {
				fmt.Println(singleIP)
			}
		}
	} else {
		hosts, err := lookIP(input)
		if err == nil {
			for _, hostname := range hosts {
				fmt.Println(hostname)
			}
		}
	}
}
