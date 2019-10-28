package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // 고루틴의 스케줄링을 확신할 수 없으므로 메인 고루틴에서 add 함수를 호출
	go func() {
		defer wg.Done() // 고루틴의 클로저를 종료하기 전에 WaitGroup에게 종료한다고 알려줌
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait() // main 고루틴은 다른 모든 고루틴이 자신들이 종료되었다고 알릴 때까지 대기
	fmt.Println("All goroutines complete")
	// Output:
	//2nd goroutine sleeping...
	//1st goroutine sleeping...
	//All goroutines complete
}
