package main

import "fmt"

func main() {
	// 슬라이스 생성
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(aSlice)
	integer := make([]int, 2)
	fmt.Println(integer)
	integer = nil
	fmt.Println(integer)
	//output:
	//[1 2 3 4 5]
	//[0 0]
	//[]

	// [:] 기호를 이용해 기존 배열을 참조하는 슬라이스 생성
	anArray := [5]int{-1, -2, -3, -4, -5}
	refAnArray := anArray[:]

	fmt.Println(anArray)
	fmt.Println(refAnArray)
	anArray[4] = -100
	fmt.Println(refAnArray)
	// output:
	//[-1 -2 -3 -4 -5]
	//[-1 -2 -3 -4 -5]
	//[-1 -2 -3 -4 -100]

	s := make([]byte, 5)
	fmt.Println(s)
	twoD := make([][]int, 3)
	fmt.Println(twoD)
	fmt.Println()
	//[0 0 0 0 0]
	//[[] [] []]

	for i := 0; i < len(twoD); i++ {
		for j := 0; j < 2; j++ {
			twoD[i] = append(twoD[i], i*j)
		}
	}

	for _, x := range twoD {
		for i, y := range x {
			fmt.Println("i:", i, "value:", y)
		}
	}
	//i: 0 value: 0
	//i: 1 value: 0
	//i: 0 value: 0
	//i: 1 value: 1
	//i: 0 value: 0
	//i: 1 value: 2
}
