package main

import (
	"fmt"
	"time"
)

func writeToChannel(c chan int, x int) {
	fmt.Println(x)
	c <- x // blocking
	close(c)
	fmt.Println(x)
}

func main() {
	c := make(chan int)
	go writeToChannel(c, 10)
	time.Sleep(1 * time.Second)
}
