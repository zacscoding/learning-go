package main

import "fmt"

func main() {
	var myInt interface{} = 123

	k, ok := myInt.(int)
	if ok {
		fmt.Println("Success:", k)
	}

	v, ok := myInt.(float64)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Failed without panicking!")
	}

	i := myInt.(int)
	fmt.Println("No checking:", i)

	j := myInt.(bool)
	fmt.Println(j)
	//Success: 123
	//Failed without panicking!
	//No checking: 123
	//panic: interface conversion: interface {} is int, not bool
	//
	//goroutine 1 [running]:
	//main.main()
	///home/zaccoding/go/src/github.com/zacscoding/learning-go/matering-go/ch7/assertion.go:23 +0x1c1
}
