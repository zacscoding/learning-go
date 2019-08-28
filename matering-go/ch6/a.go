package main

import (
	"fmt"
)

//$ mkdir ~/go/src/a
//$ cp ./a.go ~/go/src/a/
//$ go install a
func init() {
	fmt.Println("init() a")
}

func FromA() {
	fmt.Println("FromA()")
}
