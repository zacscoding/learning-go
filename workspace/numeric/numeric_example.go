package main

import (
	"fmt"
)

func main() {
	forkFactor := 0.1
	for i := 0; i <= 100; i++ {
		/*s := math.Max(float64(i)*forkFactor, 1.0)
		s = math.Min(float64(i), s)*/
		s := int(float64(i) * forkFactor)
		if s < 1 {
			s = 1
		}
		if s > i {
			s = i
		}
		fmt.Printf("Signers : %d --> %d\n", i, int(s))
	}
}
