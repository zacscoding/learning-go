package main

import (
	"errors"
	"fmt"
	"github.com/gookit/color"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	// basic1()
	// basic2()
	hasErrorTest()
}

func hasErrorTest() {
	count := 5
	aborts := make(chan struct{}, count)
	results := make(chan error, count)

	for i := 0; i < count; i++ {
		go func(idx int, r chan error, a chan struct{}) {
			s := time.Duration(rand.Intn(5)+2) * time.Second
			timer := time.NewTimer(s)
			n := rand.Intn(5)
			select {
			case <-timer.C:
				var err error
				if n == 0 {
					fmt.Println("Will push error at", idx)
					err = errors.New("n is 0 at " + strconv.Itoa(idx))
				} else {
					fmt.Println("Will push nil at", idx)
				}
				r <- err
			case <-a:
				fmt.Printf("cancel task in task-%d\n", idx)
				return
			}
		}(i, results, aborts)
	}

	for i := 0; i < count; i++ {
		res := <-results
		if res == nil {
			color.Cyan.Println("receive complete idx:", i)
		} else {
			color.Cyan.Printf("receive error in %d:%v\n", i, res)
			for j := i; j < count; j++ {
				aborts <- struct{}{}
			}
			break
		}
	}
	color.Cyan.Println("Complete tasks..")
	// Output :
	//Will push error at 2
	//Will push nil at 4
	//receive error in 0:n is 0 at 2
	//Complete tasks..
	//cancel task in task-3
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
