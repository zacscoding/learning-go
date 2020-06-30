package main

import (
	"fmt"
	"sort"
	"time"
)

type MyStruct struct {
	age int
}

const (
	TypeUnknown = iota - 1
	TypeKIP
	TypeERC
)

const (
	StatusUnknown = iota - 1
	StatusCompleted
	StatusProcessing
	StatusFailed
	StatusCancelled
)

func main() {
	// panicTest()
	// regularChannelTest()
	// testVariables()
	// testSort()
	fmt.Println("StatusUnknown:", StatusUnknown)
	fmt.Println("StatusUnknown:", StatusCompleted)
	fmt.Println("TypeUnknown:", TypeUnknown)
}

func testSort() {
	arr1 := []int{5, 3, 7, 1}
	sort.Slice(arr1, func(i, j int) bool {
		fmt.Printf("(%d, %d)\n", i, j)
		return arr1[i] > arr1[j]
	})
}

func testVariables() {
	fmt.Println(`testArgs("args1", "args2", "args3"`)
	testArgs("args1", "args2", "args3")

	var args []string
	args = append(args, "arg1")
	args = append(args, "arg2")
	args = append(args, "arg3")
	fmt.Println("testArgs with slice")
	testArgs(args...)
}

func testArgs(args ...string) {
	fmt.Println("len :", len(args))
	for _, arg := range args {
		fmt.Println(arg)
	}
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
