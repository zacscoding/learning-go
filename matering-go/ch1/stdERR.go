package main

import (
	"io"
	"os"
)

// go run stdErr.go 2>./stdErr // 표준 에러 출력을 특정 파일로 리디렉션
// go run stdErr.go 2>/dev/null // 에러 출력을 무시(/dev/null 디바이스로 리디렉션)
// go run stdErr.go > ./out 2>&1 // 표준 에러에 대한 파일 디스크립터(2)를 표준 출력에 대한 파일 디스크립터(1)로 리디렉션
// go run stdErr.go >/dev/null 2>&1

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, "This is Standard output\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
}
