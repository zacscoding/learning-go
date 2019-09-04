package main

import (
	"fmt"
	"os"
	"testing"
)

var ERR error

func benchmarkCreate(b *testing.B, buffer, filesize int) {
	var err error
	for i := 0; i < b.N; i++ {
		err = Create("/tmp/random", buffer, filesize)
	}
	ERR = err
	err = os.Remove("/tmp/random")
	if err != nil {
		fmt.Println(err)
	}
}

func Benchmark1Create(b *testing.B) {
	benchmarkCreate(b, 1, 1000000)
}

func Benchmark2Create(b *testing.B) {
	benchmarkCreate(b, 2, 1000000)
}

func Benchmark4Create(b *testing.B) {
	benchmarkCreate(b, 4, 1000000)
}

func Benchmark10Create(b *testing.B) {
	benchmarkCreate(b, 10, 1000000)
}

func Benchmark1000Create(b *testing.B) {
	benchmarkCreate(b, 1000, 1000000)
}

/*
$ go test -bench=. write_bu.go  write_bu_test.go -benchmem
goos: linux
goarch: amd64
Benchmark1Create-8                     1        1581012513 ns/op        16002808 B/op    2000023 allocs/op
Benchmark2Create-8                    10         104719180 ns/op          800433 B/op     100006 allocs/op
Benchmark4Create-8               1000000              1553 ns/op             292 B/op          5 allocs/op
Benchmark10Create-8              1000000              1308 ns/op             292 B/op          5 allocs/op
Benchmark1000Create-8            1000000              1149 ns/op             291 B/op          5 allocs/op
PASS
ok      command-line-arguments  19.178s
*/
