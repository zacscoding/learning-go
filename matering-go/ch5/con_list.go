package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func printList(l *list.List) {
	// tail ~> head
	for t := l.Back(); t != nil; t = t.Prev() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()

	// head ~> tail
	for t := l.Front(); t != nil; t = t.Next() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()
}

func main() {
	values := list.New()

	e1 := values.PushBack("One")
	e2 := values.PushBack("Two")
	// ["One", "Two"]
	values.PushFront("Three")
	// ["Three", "One", "Two"]
	values.InsertBefore("Four", e1)
	// ["Three", "Four", "One", "Two"]
	values.InsertAfter("Five", e2)
	// ["Three", "Four", "One", "Two", "Five"]
	values.Remove(e2)
	values.Remove(e2)
	// ["Three", "Four", "One", "Five"]
	values.InsertAfter("FiveFive", e2)
	// ["Three", "Four", "One", "Five"]
	values.PushBackList(values)
	// ["Three", "Four", "One", "Five", "Three", "Four", "One", "Five"]

	printList(values)
	//Five One Four Three Five One Four Three
	//Three Four One Five Three Four One Five

	values.Init()

	fmt.Printf("After Init(): %v\n", values)
	// After Init(): &{{0xc0000a2000 0xc0000a2000 <nil> <nil>} 0}

	for i := 0; i < 20; i++ {
		values.PushFront(strconv.Itoa(i))
	}

	printList(values)
	//0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19
	//19 18 17 16 15 14 13 12 11 10 9 8 7 6 5 4 3 2 1 0
}
