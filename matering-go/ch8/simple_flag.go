package main

import (
	"flag"
	"fmt"
)

func main() {
	minusK := flag.Bool("k", true, "k")
	minusO := flag.Int("O", 1, "O")
	flag.Parse()

	valueK := *minusK
	value0 := *minusO
	value0++

	fmt.Println("-k", valueK)
	fmt.Println("-0", value0)

	//$ go run simple_flag.go -O 100
	//-k true
	//-0 101
	//$ go run simple_flag.go -O=100
	//-k true
	//-0 101
	//$ go run simple_flag.go -O=100 -k false
	//-k true
	//-0 101
	//$ go run simple_flag.go -O=100 -k=false
	//-k false
	//-0 101
}
