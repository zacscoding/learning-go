package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	// 뮤텍스란 일종의 상호 배제용 락
	// 뮤텍스가 0이면 잠기지 않은 뮤텍스란 뜻
	// 한 번 사용한 뮤텍스는 복사하면 안됨
	m  sync.Mutex
	v1 int
)

func read() int {
	m.Lock()
	a := v1
	m.Unlock()
	return a
}

func change(i int) {
	m.Lock()
	time.Sleep(time.Second)
	v1 = v1 + 1
	if v1%10 == 0 {
		v1 = v1 - 10*i
	}
	m.Unlock()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give me an integer!")
		return
	}

	numGR, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var waitGroup sync.WaitGroup

	fmt.Printf("%d ", read())
	for i := 0; i < numGR; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			change(i)
			fmt.Printf("-> %d", read())
		}(i)
	}

	waitGroup.Wait()
	fmt.Printf("-> %d\n", read())
	//$ go run mutex.go 5
	//0 -> 1-> 2-> 3-> 4-> 5-> 5
}
