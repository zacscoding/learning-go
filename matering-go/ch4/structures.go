package main

import "fmt"

func main() {

	type XYZ struct {
		X int
		Y int
		Z int
	}

	var s1 XYZ
	fmt.Println(s1.Y, s1.Z)

	p1 := XYZ{23, 12, -2}
	p2 := XYZ{Z: 12, Y: 13}
	fmt.Println(p1)
	fmt.Println(p2)

	//0 0
	//{23 12 -2}
	//{0 13 12}

	pSlice := [4]XYZ{}
	pSlice[2] = p1
	pSlice[0] = p2
	fmt.Println(pSlice)
	p2 = XYZ{1, 2, 3}
	fmt.Println(pSlice)
	p1.X = -1
	fmt.Println(pSlice)
	pSlice[2].X = -50
	fmt.Println(pSlice)

	//[{0 13 12} {0 0 0} {23 12 -2} {0 0 0}]
	//[{0 13 12} {0 0 0} {23 12 -2} {0 0 0}]
	//[{0 13 12} {0 0 0} {23 12 -2} {0 0 0}]
	//[{0 13 12} {0 0 0} {-50 12 -2} {0 0 0}]

	// 포인터 리턴
	pS := new(XYZ)
	fmt.Println(pS)

	sP := new([]XYZ)
	fmt.Println(sP)

	//&{0 0 0}
	//&[]

}
