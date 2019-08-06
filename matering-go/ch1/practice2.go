package main

import (
	"fmt"
	"os"
	"strconv"
)

// 커맨드라인 인수로 입력된 실수(float) 값에 대한 평균을 구하는 프로그램
func main() {
	arguments := os.Args

	count := len(arguments) - 1

	if count == 0 {
		fmt.Println("Invalid arguments.")
		os.Exit(1)
	}

	var sum = 0.0

	for i := 1; i <= count; i++ {
		value, err := strconv.ParseFloat(arguments[i], 64)

		if err != nil {
			fmt.Println("Invalid argument :", arguments[i])
			os.Exit(1)
		}

		sum += value
	}

	average := sum / (float64(count))
	fmt.Println("Average:", average)
}
