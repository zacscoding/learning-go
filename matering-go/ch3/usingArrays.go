package main

import "fmt"

func main() {
	anArray := [4]int{1, 2, 4, -4}
	twoD := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 9}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	threeD := [2][2][2]int{{{1, 0}, {-2, 4}}, {{5, -1}, {7, 0}}}

	fmt.Println("The length of", anArray, "is", len(anArray))
	fmt.Println("The first element of", twoD, "is", twoD[0][0])
	fmt.Println("The length of", threeD, "is", len(threeD))

	isRange := true

	if (isRange) {
		for _, v := range threeD {
			for _, m := range v {
				for _, s := range m {
					fmt.Print(s, " ")
				}
			}
			fmt.Println()
		}
	} else {
		for i := 0; i < len(threeD); i++ {
			v := threeD[i]
			for j := 0; j < len(v); j++ {
				m := v[j]
				for k := 0; k < len(m); k++ {
					fmt.Print(m[k], " ")
				}
			}
		}
	}

	fmt.Println()
}
