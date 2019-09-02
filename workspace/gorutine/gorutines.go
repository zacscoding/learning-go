package main

import (
	"fmt"
	"time"
)

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

func main() {
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println()
	fmt.Println("Main Terminated")
}
