package main

import (
	"fmt"
	"sync"
	"time"
)

// 길이가 2로 고정 된 큐와 큐에 넣을 10개의 항목이 있다고 가정
func main() {
	// 표준 sync.Mutex를 Locker로 사용해 Cond 생성
	c := sync.NewCond(&sync.Mutex{}) //
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock() // 조건의 임계영역으로 들어왔기 때문에 조건과 관련된 데이터를 수정할 수 있음
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()            // 조건의 Locker에서 Lock을 호출해 조건의 임계 영역으로 진입
		for len(queue) == 2 { // 내부에서 큐의 길이를 확인
			c.Wait() // 조건에 대한 신호가 전송될 때까지 main 고루틴은 일시 중단
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second) // 큐에서 꺼내는 새로운 고루틴을 생성
		c.L.Unlock()                        // 항목을 대기열에 성공적으로 추가했으므로 조건의 임계 영역을 벗어남
	}

	// Output :
	//Adding to queue
	//Adding to queue
	//Removed from queue
	//Removed from queue
	//Adding to queue
	//Adding to queue
	//Removed from queue
	//Adding to queue
	//Removed from queue
	//Adding to queue
	//Removed from queue
	//Removed from queue
	//Adding to queue
	//Adding to queue
	//Removed from queue
	//Removed from queue
	//Adding to queue
	//Adding to queue
}
