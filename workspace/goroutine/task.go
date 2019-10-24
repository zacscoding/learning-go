package main

import "time"

type taskContext struct {
	ID       string        // task id
	Repeat   int           // repeat task
	Duration time.Duration // timeout
}

type worker interface {
	work()
}

func (t taskContext) doTask(w worker) {
	for i := 0; i < 5; i++ {
		w.work()
	}
}

func main() {
	ctx := taskContext{
		ID:       "task-1",
		Repeat:   5,
		Duration: 1 * time.Second,
	}

	ctx.doTask(worker())
}
