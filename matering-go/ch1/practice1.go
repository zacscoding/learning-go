package main

import (
	"fmt"
	"os"
	"strconv"
)

// 커맨드라인 인수로 입력된 숫자 값들을 모두 더하는 프로그램
func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Invalid arguments.")
		os.Exit(1)
	}

	sum := 0

	for i := 1; i < len(arguments); i++ {
		value, err := strconv.Atoi(arguments[i])

		if err != nil {
			fmt.Println("Invalid argument : ", arguments[i]);
			os.Exit(1)
		}

		sum += value
	}

	fmt.Println("Sum:", sum)
}
