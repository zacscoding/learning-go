package main

import "fmt"

func function1(i int) int {
	return i + i
}

func function2(i int) int {
	return i * i
}

func funFun(f func(int) int, v int) int {
	return f(v)
}

func main() {
	fmt.Println("function1:", funFun(function1, 4))
	fmt.Println("function2:", funFun(function2, 4))
	fmt.Println("Inline:", funFun(func(i int) int { return i * i * i }, 4))
	//function1: 8
	//function2: 16
	//Inline: 64
}
