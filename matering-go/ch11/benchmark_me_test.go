package main

import "testing"

var result int

func benchmarkbfibo1(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = bfibo1(n)
	}
	result = r
}

func benchmarkbfibo2(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = bfibo2(n)
	}
	result = r
}

func benchmarkbfibo3(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = bfibo3(n)
	}
	result = r
}

func Benchmark30bfibo1(b *testing.B) {
	benchmarkbfibo1(b, 30)
}

func Benchmark30bfibo2(b *testing.B) {
	benchmarkbfibo2(b, 30)
}

func Benchmark30bfibo3(b *testing.B) {
	benchmarkbfibo3(b, 30)
}

func Benchmark50bfibo1(b *testing.B) {
	benchmarkbfibo1(b, 50)
}

func Benchmark50bfibo2(b *testing.B) {
	benchmarkbfibo2(b, 50)
}

func Benchmark50bfibo3(b *testing.B) {
	benchmarkbfibo3(b, 50)
}

/*
$ go test -bench=. benchmark_me.go benchmark_me_test.go
goos: linux
goarch: amd64
Benchmark30bfibo1-8          300           4736850 ns/op
Benchmark30bfibo2-8          300           6267284 ns/op
Benchmark30bfibo3-8       300000              3988 ns/op
Benchmark50bfibo1-8            1        82683185836 ns/op
Benchmark50bfibo2-8            1        85306563027 ns/op
Benchmark50bfibo3-8       300000              5123 ns/op
PASS
ok      command-line-arguments  175.123s
*/
