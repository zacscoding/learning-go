package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need a domain name!")
		return
	}

	domain := arguments[1]
	MXs, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, MX := range MXs {
		fmt.Println(MX.Host)
	}
}

//$ go run mx_records.go google.com
//aspmx.l.google.com.
//alt1.aspmx.l.google.com.
//alt2.aspmx.l.google.com.
//alt3.aspmx.l.google.com.
//alt4.aspmx.l.google.com.
//
//$ host -t mx google.com
//google.com mail is handled by 30 alt2.aspmx.l.google.com.
//google.com mail is handled by 40 alt3.aspmx.l.google.com.
//google.com mail is handled by 10 aspmx.l.google.com.
//google.com mail is handled by 20 alt1.aspmx.l.google.com.
//google.com mail is handled by 50 alt4.aspmx.l.google.com.
