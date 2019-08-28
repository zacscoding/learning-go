package main

import (
	"container/ring"
	"fmt"
)

var size int = 10

func main() {
	myRing := ring.New(size + 1)
	fmt.Println("Empty ring:", *myRing)
	//Empty ring: {0xc00000c060 0xc00000c180 <nil>}

	for i := 0; i < myRing.Len()-1; i++ {
		myRing.Value = i
		myRing = myRing.Next()
	}

	myRing.Value = 2

	sum := 0
	myRing.Do(func(x interface{}) {
		t := x.(int) // type assertion
		fmt.Print(t, " ")
		sum += t
	})
	fmt.Println()
	fmt.Println("Sum:", sum)
	//2 0 1 2 3 4 5 6 7 8 9
	//Sum: 47

	for i := 0; i < myRing.Len()+2; i++ {
		myRing = myRing.Next()
		fmt.Print(myRing.Value, " ")
	}
	fmt.Println()
	// 0 1 2 3 4 5 6 7 8 9 2 0 1
}
