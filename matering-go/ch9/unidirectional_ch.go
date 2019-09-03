package main

import (
	"fmt"
	"time"
)

func writeOnly(c chan<- int, x int) {
	fmt.Println("Write to :", x)
	c <- x
	// a := <-c // compile error
	close(c)
}

func readOnly(c <-chan int) {
	// c <- 10 // compile err
	fmt.Println("Read :", <-c)
}

func main() {
	c := make(chan int)
	go readOnly(c)
	go writeOnly(c, 10)
	time.Sleep(1 * time.Second)
}
