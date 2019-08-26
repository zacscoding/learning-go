package main

import "fmt"

func getPointer(n *int) {
	// n이 가리키는 값 제곱
	*n = *n * *n
}

func returnPointer(n int) *int {
	v := n * n
	return &v
}

func main() {
	i := -10
	j := 25

	pI := &i
	pJ := &j

	fmt.Println("pI memory:", pI)
	fmt.Println("pJ memory:", pJ)
	fmt.Println("pI value:", *pI)
	fmt.Println("pJ value:", *pJ)
	//pI memory: 0xc000098000
	//pJ memory: 0xc000098008
	//pI value: -10
	//pJ value: 25

	*pI = 123456
	*pI--
	fmt.Println("i:", i)
	// i: 123455

	getPointer(pJ)
	fmt.Println("j:", j)
	k := returnPointer(12)
	fmt.Println(*k)
	fmt.Println(k)
	//j: 625
	//144
	//0xc000098050

}
