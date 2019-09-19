package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	var m sync.Mutex
	count := 0
	timeout := time.NewTimer(5 * time.Second)
	complete := make(chan bool)
	defer timeout.Stop()

	for i := 0; i < 100; i++ {
		go func(k int) {
			if rand.Intn(100)%30 == 0 {
				m.Lock()
				count++
				m.Unlock()
				if count == 5 {
					complete <- true
				}
			}
			time.Sleep(100 * time.Millisecond)
		}(i)
	}

	select {
	case <-complete:
		fmt.Println("Complete!!")
	case <-timeout.C:
		fmt.Println("timeout!!")
	}
}
