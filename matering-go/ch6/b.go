package main

import (
	"a"
	"fmt"
)

//$ mkdir ~/go/src/b
//$ cp ./b.go ~/go/src/b/
//$ go install b
func init() {
	fmt.Println("init() b")
}

func FromB() {
	fmt.Println("FromB()")
	a.FromA()
}
