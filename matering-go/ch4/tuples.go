package main

import "fmt"

// 세 개의 정수를 가진 튜플 반환
func retThree(x int) (int, int, int) {
	return 2 * x, x * x, -x
}

func main() {
	fmt.Println(retThree(10))
	// 튜플 할당
	n1, n2, n3 := retThree(20)
	fmt.Println(n1, n2, n3)
	//20 100 -10
	//40 400 -20

	// swap
	n1, n2 = n2, n1
	fmt.Println(n1, n2, n3)
	x1, x2, x3 := n1*2, n1*n1, -n1
	fmt.Println(x1, x2, x3)

	//400 40 -20
	//800 160000 -400
}
