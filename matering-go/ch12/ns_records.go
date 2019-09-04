package main

import (
	"fmt"
	"net"
	"os"
)

// 도메인의 네임 서버를 찾는 예제
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need a domain name!")
		return
	}

	domain := arguments[1]
	NSs, err := net.LookupNS(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, NS := range NSs {
		fmt.Println(NS.Host)
	}
}

//$ go run ns_records.go google.com
//ns3.google.com.
//ns2.google.com.
//ns4.google.com.
//ns1.google.com.
