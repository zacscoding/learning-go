package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)

	for {
		if c == nil {
			fmt.Println("for loop executed after c == nil")
		}
		// nil 채널 블로킹
		select {
		case input := <-c:
			sum += input
		case <-t.C: // 지정한 시간만큼 t 타이머의 C채널을 블록
			c = nil
			fmt.Println(sum)
		}
	}
}

func send(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}

func main() {
	c := make(chan int)
	go add(c)
	go send(c)

	time.Sleep(5 * time.Second)
}
