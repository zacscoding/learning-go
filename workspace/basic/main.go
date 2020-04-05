package main

import (
	"fmt"
	"time"
)

type MyStruct struct {
	age int
}

func main() {
	// panicTest()
	regularChannelTest()
}

func regularChannelTest() {
	myChannel := make(chan int)
	// myBufferedChannel := make(chan int, 10) // 10개의 int 버퍼 채널
	go runLoopSend(10, myChannel)
	go runLoopReceive(myChannel)

	time.Sleep(2 * time.Second)
}

func runLoopSend(n int, ch chan int) {
	for i := 0; i < n; i++ {
		ch <- i
	}
	close(ch)
}

func runLoopReceive(ch chan int) {
	for {
		// 채널이 열려 있으면 ok == true, 닫혀 있으면 ok == false
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Receive value:", i)
	}

	//for i := range ch {
	//	fmt.Println("Receive value:", i)
	//}
}

func panicTest() {
	forcePanic(true)
	fmt.Println("Hello world")
}

func checkPanic() {
	if r := recover(); r != nil {
		fmt.Println("A Panic was captured, message :", r)
	}
}

func forcePanic(p bool) {
	defer checkPanic()
	if p {
		panic("panic requested")
	}
}

func callByTest() {
	s1 := MyStruct{
		age: 10,
	}
	fmt.Printf("made struct : %p\n", &s1)
	callByRef(s1)
}

func callByPointer(s *MyStruct) {
	fmt.Printf("callByPointer is called.. %s", s)
}

func callByRef(s MyStruct) {
	fmt.Printf("callByRef : %p\n", &s)
}
