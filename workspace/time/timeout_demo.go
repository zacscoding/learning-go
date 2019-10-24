package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	// timeout1()
	// timeout2()
	timeout3()
}

type task func(taskID string, count int)

func timeout3() {
	// given
	taskCount := 3
	repeat := 0
	duration := 5 * time.Second
	interval := 1 * time.Second

	task := func(taskID string, count int) {
		fmt.Printf("[%s] working #%d\n", taskID, count)
	}

	var wait sync.WaitGroup
	var cancel chan bool
	for i := 0; i < taskCount; i++ {
		wait.Add(1)
		taskID := "task-" + strconv.Itoa(i+1)
		c := schedule(taskID, repeat, interval, duration, &wait, task)
		if i == 0 {
			cancel = c
		}
	}

	go func() {
		time.Sleep(4 * time.Second)
		cancel <- true
	}()

	wait.Wait()
	fmt.Println("## Complete task all")
}

func schedule(taskID string, repeat int, interval, delay time.Duration, w *sync.WaitGroup, t task) chan bool {
	fmt.Printf("[%s] > delay : %s\n", taskID, delay)
	ticker := time.NewTicker(interval)
	cancel := make(chan bool)
	timeout := time.NewTimer(delay)
	go func() {
		idx := 1
		defer w.Done()
		for {
			if repeat > 0 && idx > repeat {
				return
			}
			select {
			case <-timeout.C:
				fmt.Printf("[%s] timeout at %d\n", taskID, idx)
				return
			case <-cancel:
				fmt.Printf("[%s] cancel at %d\n", taskID, idx)
				return
			case <-ticker.C:
				t(taskID, idx)
				idx++
			}
		}
	}()
	return cancel
}

func timeout2() {
	// given
	taskCount := 5
	repeat := 0
	timeout := 5 * time.Second

	if repeat > 0 {
		var wait sync.WaitGroup
		for i := 0; i < taskCount; i++ {
			wait.Add(1)
			go func(idx int, w *sync.WaitGroup) {
				timeout2Internal(idx, repeat, w, nil)
			}(i, &wait)
		}
		wait.Wait()
	} else {
		timer := time.NewTimer(timeout)
		defer timer.Stop()
		quits := make([]chan bool, taskCount)
		for i := 0; i < taskCount; i++ {
			quits[i] = make(chan bool)
			go func(idx int, q chan bool) {
				timeout2Internal(idx, repeat, nil, q)
			}(i, quits[i])
		}
		select {
		case <-timer.C:
			for _, quit := range quits {
				go func() {
					quit <- true
				}()
			}
		}
	}
	fmt.Println("### Complete all task!!")
}

func timeout2Internal(id int, repeat int, waitGroup *sync.WaitGroup, quit chan bool) {
	if repeat > 0 {
		for i := 0; i < repeat; i++ {
			fmt.Printf("Do work %d in task-%d\n", i, id)
			time.Sleep(1 * time.Second)
		}
		waitGroup.Done()
		return
	}

	idx := 0
	for {
		select {
		case <-quit:
			fmt.Printf("Terminate go routine task-%d\n", id)
			return
		default:
			fmt.Printf("Do work %d in task-%d\n", idx, id)
			idx++
		}
		time.Sleep(1 * time.Second)
	}
}

func timeout1() {
	var m sync.Mutex
	count := 0
	timeout := time.NewTimer(5 * time.Second)
	complete := make(chan bool)
	defer timeout.Stop()

	for i := 0; i < 100; i++ {
		go func(k int) {
			if rand.Intn(100)%30 == 0 {
				m.Lock()
				count++
				m.Unlock()
				if count == 5 {
					complete <- true
				}
			}
			time.Sleep(100 * time.Millisecond)
		}(i)
	}

	select {
	case <-complete:
		fmt.Println("Complete!!")
	case <-timeout.C:
		fmt.Println("timeout!!")
	}
}
