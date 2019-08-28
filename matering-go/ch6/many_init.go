package main

import (
	"a"
	"b"
	"fmt"
)

func init() {
	fmt.Println("init() manyInit")
}

func main() {
	a.FromA()
	b.FromB()
	//init() a
	//init() b
	//init() manyInit
	//FromA()
	//FromB()
	//FromA()
}
