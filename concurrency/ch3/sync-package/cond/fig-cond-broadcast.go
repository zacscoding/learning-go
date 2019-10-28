package main

import (
	"fmt"
	"sync"
)

func main() {
	// Clicked라는 조건을 가지고 있는 Button 타입 지정
	type Button struct {
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	// 조건의 신호들을 처리하는 함수를 등록할 수 있는 편의함수를 제공
	// 각 핸들러는 자체 고루틴에서 실행되며, 고루틴이 실행 중이라는 것을 확인하기 전까지
	// subscribe 종료되지 않는다.
	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	// 3개의 핸들러 등록
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast() // 이벤트 트리거
	clickRegistered.Wait()
	// Output :
	//Mouse clicked.
	//Maximizing window.
	//Displaying annoying dialog box!
}
