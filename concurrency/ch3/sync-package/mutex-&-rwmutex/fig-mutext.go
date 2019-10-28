package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}
	arithmetic.Wait()
	fmt.Println("Arithmetic complete. > ", count)

	// Output :
	//Decrementing: -1
	//Incrementing: 0
	//Incrementing: 1
	//Incrementing: 2
	//Decrementing: 1
	//Decrementing: 0
	//Decrementing: -1
	//Incrementing: 0
	//Decrementing: -1
	//Decrementing: -2
	//Incrementing: -1
	//Incrementing: 0
	//Arithmetic complete. >  0
}
