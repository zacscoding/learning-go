package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func fibo1(n int) int64 {
	if n == 0 || n == 1 {
		return int64(n)
	}
	time.Sleep(time.Millisecond)
	return int64(fibo2(n-1)) + int64(fibo2(n-2))
}

func fibo2(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	time.Sleep(50 * time.Millisecond)
	return fn[n]
}

// check primary
func N1(n int) bool {
	k := math.Floor(float64(n/2 + 1))
	for i := 2; i < int(k); i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

func N2(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	cpuFile, err := os.Create("/tmp/cpuProfile.out")
	if err != nil {
		fmt.Println(err)
		return
	}

	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()

	total := 0
	for i := 2; i < 100000; i++ {
		n := N1(i)
		if n {
			total++
		}
	}
	fmt.Println("Total primes:", total)

	total = 0
	for i := 2; i < 100000; i++ {
		n := N2(i)
		if n {
			total++
		}
	}
	fmt.Println("Total primes:", total)

	for i := 1; i < 90; i++ {
		n := fibo1(i)
		fmt.Print(n, " ")
	}
	fmt.Println()

	for i := 1; i < 90; i++ {
		n := fibo2(i)
		fmt.Print(n, " ")
	}
	fmt.Println()
	runtime.GC()

	// 메모리 프로파일링
	memory, err := os.Create("/tmp/memoryProfile.out")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer memory.Close()
	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(50 * time.Millisecond)
	}
	err = pprof.WriteHeapProfile(memory)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//$ go tool pprof /tmp/cpuProfile.out
//$ go tool pprof -http=localhost:8080 /tmp/cpuProfile.out

// (pprof) top
// (pprof) top10 --cum
// (pprof) list main.N1
// (pprof) pdf


// (pprof) top
//File: ___go_build_github_com_zacscoding_learning_go_matering_go_ch11
//Type: cpu
//Time: Sep 4, 2019 at 11:13am (KST)
//Duration: 20.05s, Total samples = 6.13s (30.58%)
//Entering interactive mode (type "help" for commands, "o" for options)
//(pprof) top
//Showing nodes accounting for 6.09s, 99.35% of 6.13s total
//Dropped 9 nodes (cum <= 0.03s)
//Showing top 10 nodes out of 11
//flat  flat%   sum%        cum   cum%
//3.84s 62.64% 62.64%      3.84s 62.64%  main.N2
//2.20s 35.89% 98.53%      2.20s 35.89%  main.N1
//0.05s  0.82% 99.35%      0.05s  0.82%  runtime.memclrNoHeapPointers
//0     0% 99.35%      6.12s 99.84%  main.main
//0     0% 99.35%      0.05s  0.82%  runtime.(*mheap).alloc
//0     0% 99.35%      0.05s  0.82%  runtime.largeAlloc
//0     0% 99.35%      6.12s 99.84%  runtime.main
//0     0% 99.35%      0.05s  0.82%  runtime.makeslice
//0     0% 99.35%      0.05s  0.82%  runtime.mallocgc
//0     0% 99.35%      0.05s  0.82%  runtime.mallocgc.func1
