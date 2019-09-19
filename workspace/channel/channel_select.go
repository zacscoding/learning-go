package channelmulti

import (
	"fmt"
	"math/rand"
	"time"
)

func write(ch chan int) {
	for {
		r := rand.Intn(10)
		ch <- r
		time.Sleep(1 * time.Second)
	}
}

func terminate(ch chan bool) {
	for {
		r := rand.Intn(30)
		if r == 29 {
			ch <- true
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan bool)

	go write(ch1)
	go terminate(ch2)

EXIT:
	for {
		select {
		case read := <-ch1:
			fmt.Println("Read...", read)
		case <-ch2:
			fmt.Println("Will terminate")
			break EXIT
		}
	}
	fmt.Println("Complete!")
}
