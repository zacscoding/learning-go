package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// STOP 이란 단어를 입력할 때까지 계속해서 입력된 정수 값을 읽는 프로그램
func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		read := scanner.Text()

		if read == "STOP" {
			fmt.Println("Will terminate this program :)")
			break
		}

		integer, err := strconv.Atoi(read)

		if err != nil {
			fmt.Printf("Cannot cast \"%s\" to integer", read)
			fmt.Println()
		} else {
			fmt.Println("Read:", integer)
		}
	}
}
