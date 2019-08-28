package main

import "fmt"

func funReturnFun() func() int {
	i := 0
	// 익명 함수 리턴
	return func() int {
		i++
		return i * i
	}
}

func main() {
	i := funReturnFun()
	j := funReturnFun()

	fmt.Println("1:", i())
	fmt.Println("2:", i())

	fmt.Println("j1:", j())
	fmt.Println("j2:", j())

	fmt.Println("3:", i())
	//1: 1
	//2: 4
	//j1: 1
	//j2: 4
	//3: 9
}
