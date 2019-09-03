package main

import (
	"fmt"
	"time"
)

func A(a, b chan struct{}) {
	<-a
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(b)
}

func B(a, b chan struct{}) {
	<-a
	fmt.Println("B()!")
	close(b)
}

func C(a chan struct{}) {
	<-a
	fmt.Println("C()!")
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	// C() 에서 z 채널로 블로킹
	go C(z)
	// A() 에서 x 채널로 블로킹 후 y 채널 close
	go A(x, y)
	// C() 에서 z 채널 블로킹
	go C(z)
	// B()에서 y채널로 블로킹 후 z 채널 close
	go B(y, z)
	// C() 에서 z 채널 블로킹
	go C(z)

	close(x)
	time.Sleep(3 * time.Second)
	//A()!
	//B()!
	//C()!
	//C()!
	//C()!
}
