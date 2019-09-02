package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(signal os.Signal) {
	fmt.Println("handleSignal() Caught:", signal)
}

// $ go build hello_two.go
// $ ./hello_two
// $ ps ax | grep ./hello_two | grep -v grep
//18688 pts/0    Tl     0:00 ./hello_two
//18718 pts/0    Sl+    0:00 ./hello_two
func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGKILL)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				fmt.Println("Caught:", sig)
			case syscall.SIGKILL:
				handleSignal(sig)
			}
		}
	}()

	for {
		fmt.Printf(".")
		time.Sleep(20 * time.Second)
	}

}
