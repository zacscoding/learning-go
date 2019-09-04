package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("You are using", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("with Go version", runtime.Version())
}

//You are usinggc on a amd64 machine
//with Go version go1.12.1

//$ env GOOS=linux GOARCH=arm go build x_compile.go
//$ ./x_compile
//bash: ./x_compile: cannot execute binary file: Exec 형식 오류
//$ file x_compile
//x_compile: ELF 32-bit LSB executable, ARM, EABI5 version 1 (SYSV), statically linked, not stripped
