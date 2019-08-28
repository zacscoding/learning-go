package main

import (
	"fmt"
	"os"
	"strconv"
)

func doubleSquare(x int) (int, int) {
	return x * 2, x * x
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("The program needs 1 arguments!")
		return
	}

	y, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	square := func(s int) int {
		return s * s
	}
	fmt.Println("The square of", y, "is", square(y))
	//The square of  3 is 9

	double := func(s int) int {
		return 2 * s
	}
	fmt.Println("The double of", y, "is", double(y))
	//The double of  3 is 6

	fmt.Println(doubleSquare(y))
	//6 9

	d, s := doubleSquare(y)
	fmt.Println(d, s)
	//6 9
}
