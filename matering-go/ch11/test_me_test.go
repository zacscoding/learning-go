package main

import "testing"

/*
$ go test test_me.go test_me_test.go -v
=== RUN   TestS1
--- FAIL: TestS1 (0.00s)
test_me_test.go:7: s1("123456789") != 9"
=== RUN   TestS2
--- PASS: TestS2 (0.00s)
=== RUN   TestF1
--- FAIL: TestF1 (0.00s)
test_me_test.go:33: f1(2) != 2
=== RUN   TestF2
--- FAIL: TestF2 (0.00s)
test_me_test.go:45: f2(1) != 1
test_me_test.go:51: f2(10) != 55
FAIL
FAIL    command-line-arguments  0.001s
*/

func TestS1(t *testing.T) {
	if s1("123456789") != 9 {
		t.Error(`s1("123456789") != 9"`)
	}

	if s1("") != 0 {
		t.Error(`s1("") != 0`)
	}
}

func TestS2(t *testing.T) {
	if s2("123456789") != 9 {
		t.Error(`s2("123456789") != 9"`)
	}

	if s2("") != 0 {
		t.Error(`s2("") != 0`)
	}
}

func TestF1(t *testing.T) {
	if f1(0) != 0 {
		t.Error(`f1(0) != 0`)
	}
	if f1(1) != 1 {
		t.Error(`f1(1) != 1`)
	}
	if f1(2) != 2 {
		t.Error(`f1(2) != 2`)
	}
	if f1(10) != 55 {
		t.Error(`f1(10) != 55`)
	}
}

func TestF2(t *testing.T) {
	if f2(0) != 0 {
		t.Error(`f2(0) != 0`)
	}
	if f2(1) != 1 {
		t.Error(`f2(1) != 1`)
	}
	if f2(2) != 2 {
		t.Error(`f2(2) != 2`)
	}
	if f2(10) != 55 {
		t.Error(`f2(10) != 55`)
	}
}
