package main

import (
	"fmt"
	"aPackage"
)

//$ mkdir ~/go/src/aPackage
//$ cp aPackage/a_package.go ~/go/src/aPackage/
//$ go install aPackage
//$ ls -l ~/go/pkg/linux_amd64/aPackage.a
//-rw-rw-r-- 1 zaccoding zaccoding 16564  8월 28 16:19 /home/zaccoding/go/pkg/linux_amd64/aPackage.a
func main() {
	fmt.Println("Using aPackage!")
	aPackage.A()
	aPackage.B()
	fmt.Println(aPackage.MyConstant)
	//Using aPackage!
	//This is function A!
	//privateConstant: 21
	//123
}
