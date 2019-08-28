package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var size = 0
var queue = new(Node)

func Push(t *Node, v int) bool {
	if queue == nil {
		queue = &Node{v, nil}
		size++
		return true
	}

	t = &Node{v, nil}
	t.Next = queue
	queue = t
	size++

	return true
}

func Pop(t *Node) (int, bool) {
	if size == 0 {
		return 0, false
	}

	if size == 1 {
		queue = nil
		size--
		return t.Value, true
	}

	temp := t
	for (t.Next) != nil {
		temp = t
		t = t.Next
	}

	v := (temp.Next).Value
	temp.Next = nil

	size--
	return v, true
}

func traverse(t *Node) {
	if size == 0 {
		fmt.Println("Empty Queue!")
		return
	}

	for t != nil {
		fmt.Print(t.Value)
		t = t.Next
		if t != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

func main() {
	queue = nil
	Push(queue, 10)
	fmt.Println("Size:", size)
	traverse(queue)
	//Size: 1
	//10

	v, b := Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)
	//Pop: 10
	//Size: 0

	for i := 0; i < 5; i++ {
		Push(queue, i)
	}
	traverse(queue)
	fmt.Println("Size:", size)
	//4 -> 3 -> 2 -> 1 -> 0
	//Size: 5

	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)
	//Pop: 0
	//Size: 4

	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)
	traverse(queue)
	//Pop: 1
	//Size: 3
	//4 -> 3 -> 2
}
