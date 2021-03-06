package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// $ sudo lsof -n -i :8001
// COMMAND    PID      USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
// other_tcp 2321 zaccoding    3u  IPv4 631078      0t0  TCP 127.0.0.1:8001 (LISTEN)
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	SERVER := "127.0.0.1" + ":" + arguments[1]

	s, err := net.ResolveTCPAddr("tcp", SERVER)
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := net.ListenTCP("tcp", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := make([]byte, 1024)
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting TCP server!")
			conn.Close()
			return
		}

		fmt.Print("> ", string(buffer[0:n-1]))
		_, err = conn.Write(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
