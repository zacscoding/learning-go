package ex

import "fmt"

func ExampleF1() {
	fmt.Println(F1(10))
	fmt.Println(F1(2))
	// Output:
	// 55
	// 1
}

func ExampleS1() {
	fmt.Println(S1("123456789"))
	fmt.Println(S1(""))
	// Output:
	// 8
	// 0
}

/*
$ go test ex.go ex_test.go -v
== = RUN   ExampleF1
--- PASS: ExampleF1 (0.00s)
== = RUN   ExampleS1
--- FAIL: ExampleS1 (0.00s)
got:
9
0
want:
8
0
FAIL
FAIL    command-line-arguments  0.002s
*/
