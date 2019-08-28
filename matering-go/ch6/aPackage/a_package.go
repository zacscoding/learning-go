package aPackage

import "fmt"

const MyConstant = 123
const privateConstant = 21

func A() {
	fmt.Println("This is function A!")
}

func B() {
	fmt.Println("privateConstant:", privateConstant)
}
