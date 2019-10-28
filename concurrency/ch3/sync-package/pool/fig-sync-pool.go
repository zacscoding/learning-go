package main

import (
	"fmt"
	"sync"
)

func main() {
	var numCalcCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	// 4KB 풀로 시작
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcCreated)

	// Output :
	// 8 calculators were created.
}
