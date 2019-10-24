package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// basic1()
	basic2()
}

func basic1() {
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println()
	fmt.Println("terminated basic1()")
	// Output:
	//0----250ms----400ms----500ms----
	//------|--------|--------|
	//------1--------a--------2
}

// 5 go routines and wait for completion and collect data
func basic2() {
	count := 5
	var wait sync.WaitGroup
	var results []int
	ch := make(chan int)

	for i := 0; i < count; i++ {
		wait.Add(1)
		go func(ch chan int, idx int) {
			defer wait.Done()
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			fmt.Println("will write :", idx)
			ch <- idx
		}(ch, i+1)
	}

	go func(results *[]int) {
		for c := range ch {
			fmt.Println("Receive : ", c)
		}
	}(&results)
	wait.Wait()
	fmt.Println("terminate")

	// Output :
	//will write : 5
	//Receive :  5
	//will write : 1
	//Receive :  1
	//will write : 3
	//Receive :  3
	//will write : 2
	//Receive :  2
	//will write : 4
	//Receive :  4
	//terminate
}

func numbers() {
	for i := 1; i < 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c", i)
	}
}

/**
0----250ms----400ms----500ms----
------|--------|--------|
------1--------a--------2
*/
