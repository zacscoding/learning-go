package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func startServer(port string) {
	go func(port string) {
		http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			time.Sleep(2 * time.Second)
			fmt.Fprintf(writer, "Success")
		})
		err := http.ListenAndServe(port, nil)
		if err != nil {
			fmt.Println(err)
		}
	}(port)
}

var timeout = time.Duration(time.Second)

func Timeout(network, host string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, host, timeout)
	if err != nil {
		return nil, err
	}
	conn.SetDeadline(time.Now().Add(timeout))
	return conn, nil
}

//$ go run client_timeout.go http://localhost:8081 1
//1  gone..
//Get http://localhost:8081: read tcp 127.0.0.1:49020->127.0.0.1:8081: i/o timeout
func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage %s URL TIMEOUT\n", filepath.Base(os.Args[0]))
		return
	}

	if len(os.Args) == 3 {
		temp, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Using default timeout")
		} else {
			timeout = time.Duration(time.Duration(temp) * time.Second)
		}
	}

	startServer(":8081")

	URL := os.Args[1]
	t := http.Transport{
		Dial: Timeout,
	}

	client := http.Client{
		Transport: &t,
	}

	go func() {
		var waitSeconds = 1
		for {
			time.Sleep(time.Second)
			fmt.Println(waitSeconds, " gone..")
			waitSeconds++
		}
	}()
	data, err := client.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer data.Body.Close()
	_, err = io.Copy(os.Stdout, data.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

}
