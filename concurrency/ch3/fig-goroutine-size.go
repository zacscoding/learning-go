package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c chan interface{}
	var wg sync.WaitGroup
	// 프로세스가 끝날때까지 종료X
	noop := func() {
		wg.Done()
		<-c
	}

	const numGoroutines = 1e4 // 고루틴 수
	wg.Add(numGoroutines)
	before := memConsumed() // 고루틴들을 생성하기 전 소비 된 메모리 양
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed() // 고루틴들을 생성 한 후 소비 된 메모리 양
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
	// Output:
	// 0.013kb
}
