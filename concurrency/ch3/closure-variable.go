package main

import (
	"fmt"
	"sync"
)

func main() {
	closureWithVariable()
}

func closureWithVariable() {
	fmt.Println("// test1")
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("salutation:", salutation)
		}()
	}
	wg.Wait()
	fmt.Println("// test2")
	for i, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		fmt.Printf("in for loop #%d &salutation: %v\n", i, &salutation)
		go func(i int, salutation string) { // shadow variable
			defer wg.Done()
			fmt.Printf("in goroutine #%d &salutation: %v\n", i, &salutation)
			fmt.Println("salutation:", salutation)
		}(i, salutation)
	}
	wg.Wait()
	// Output:
	//// test1
	//salutation: good day
	//salutation: good day
	//salutation: good day
	//// test2
	//in for loop #0 &salutation: 0xc000010220
	//in for loop #1 &salutation: 0xc000010220
	//in for loop #2 &salutation: 0xc000010220
	//in goroutine #2 &salutation: 0xc000010230
	//salutation: good day
	//in goroutine #0 &salutation: 0xc000010250
	//salutation: hello
	//in goroutine #1 &salutation: 0xc000096000
	//salutation: greetings0
}
