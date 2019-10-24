package main

import "fmt"

type MyStruct struct {
	age int
}

func main() {
}

func callByTest() {
	s1 := MyStruct{
		age: 10,
	}
	fmt.Printf("made struct : %p\n", &s1)
	callByRef(s1)
}

func callByPointer(s *MyStruct) {
	fmt.Printf("callByPointer is called.. %s", s)
}

func callByRef(s MyStruct) {
	fmt.Printf("callByRef : %p\n", &s)
}
