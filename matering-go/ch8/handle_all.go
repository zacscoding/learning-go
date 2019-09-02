package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handle(signal os.Signal) {
	fmt.Println("Received:", signal)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)

	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				handle(sig)
			case syscall.SIGTERM:
				handle(sig)
				os.Exit(0)
			case syscall.SIGUSR2:
				fmt.Println("Handling syscall.SIGUSR2")
			default:
				fmt.Println("Ignoring", sig)
			}
		}
	}()

	for {
		fmt.Printf(".")
		time.Sleep(5 * time.Second)
	}

	//$ ps ax | grep ./handle_all | grep -v grep
	//19022 pts/1    Sl+    0:00 ./handle_all
	//$ kill -s HUP 19022
	//$ kill -s USR2 19022
	//$ kill -s USR1 19022
	//$ kill -s int 19022
	//$ kill -s term 19022
}
