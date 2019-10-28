package main

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		// 고루틴을 설정하고 시작하는 비용이 컨텍스트 스위칭을 측정하는데 영향을 미치길 원하지X
		<-begin
		for i := 0; i < b.N; i++ {
			// 빈 구조체로 수신측 고루틴에게 메시지를 보냄
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		// 고루틴을 설정하고 시작하는 비용이 컨텍스트 스위칭을 측정하는데 영향을 미치길 원하지X
		<-begin
		for i := 0; i < b.N; i++ {
			// 메시지를 수신하지만 아무런 작업X
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	// 성능 타이머 시작
	b.StartTimer()
	close(begin)
	wg.Wait()

	// Output :
	//$ go test -bench=. -cpu=4 fig-ctx-switch_test.go
	//goos: linux
	//goarch: amd64
	//BenchmarkContextSwitch-4        10000000               155 ns/op
	//PASS
	//ok      command-line-arguments  1.715s
}
