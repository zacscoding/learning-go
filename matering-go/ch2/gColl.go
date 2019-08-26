package main

import (
	"fmt"
	"runtime"
	"time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)

	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}

func main() {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}

		time.Sleep(5 * time.Second)
	}

	printStats(mem)
}

//GODEBUG=gctrace=1 go run gColl.go
//gc 1 @0.041s 0%: 0.062+0.69+0.009 ms clock, 0.49+0.23/0.82/0.62+0.073 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
//gc 2 @0.052s 0%: 0.033+0.32+0.006 ms clock, 0.27+0.13/0.37/0.84+0.053 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
//gc 3 @0.059s 1%: 0.13+5.0+0.014 ms clock, 1.1+0.77/0.86/0+0.11 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
//gc 4 @0.067s 1%: 0.005+0.52+0.005 ms clock, 0.045+0.35/0.57/0.76+0.040 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
//gc 5 @0.069s 1%: 0.004+0.68+0.004 ms clock, 0.033+0.043/0.36/0.96+0.038 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
//gc 6 @0.074s 1%: 0.003+0.41+0.003 ms clock, 0.024+0.084/0.51/0.64+0.026 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
//gc 7 @0.077s 1%: 0.003+0.35+0.006 ms clock, 0.026+0.085/0.49/0.64+0.053 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
//# command-line-arguments
//gc 1 @0.001s 10%: 0.002+1.6+0.010 ms clock, 0.018+0.56/1.8/0.17+0.081 ms cpu, 5->6->6 MB, 6 MB goal, 8 P
//gc 2 @0.007s 7%: 0.005+2.0+0.008 ms clock, 0.041+0.096/2.8/0.69+0.065 ms cpu, 8->8->8 MB, 12 MB goal, 8 P
//gc 3 @0.013s 6%: 0.004+3.8+0.004 ms clock, 0.035+0.22/3.5/2.6+0.037 ms cpu, 15->16->15 MB, 16 MB goal, 8 P
//gc 4 @0.040s 4%: 0.004+9.8+0.002 ms clock, 0.035+0.26/9.5/0.17+0.021 ms cpu, 25->25->24 MB, 31 MB goal, 8 P
//mem.Alloc: 109000
//mem.TotalAlloc: 109000
//mem.HeapAlloc 109000
//mem.NumGC: 0
//-----
//gc 1 @0.001s 4%: 0.050+0.10+0.002 ms clock, 0.40+0.041/0.023/0.059+0.023 ms cpu, 47->47->0 MB, 48 MB goal, 8 P
//gc 2 @0.022s 0%: 0.002+0.10+0.002 ms clock, 0.021+0.062/0.058/0.034+0.023 ms cpu, 47->47->0 MB, 48 MB goal, 8 P
//gc 3 @0.024s 0%: 0.001+0.085+0.002 ms clock, 0.014+0.059/0/0.017+0.019 ms cpu, 47->47->0 MB, 48 MB goal, 8 P
//gc 4 @0.026s 0%: 0.002+0.11+0.002 ms clock, 0.016+0.069/0.020/0.067+0.022 ms cpu, 47->47->0 MB, 48 MB goal, 8 P
//gc 5 @0.029s 0%: 0.001+0.093+0.003 ms clock, 0.010+0.075/0.015/0.074+0.025 ms cpu, 47->47->0 MB, 48 MB goal, 8 P
//gc 6 @0.031s 0%: 0.002+0.094+0.003 ms clock, 0.021+0.079/0.004/0.064+0.024 ms cpu, 47->47->0 MB, 48 MB goal, 8 P
//gc 7 @0.034s 0%: 0.002+0.068+0.002 ms clock, 0.020+0.058/0.029/0.041+0.022 ms cpu, 47->47->0 MB, 48 MB goal, 8 P
//gc 8 @0.036s 0%: 0.001+0.090+0.002 ms clock, 0.015+0.076/0.004/0.056+0.018 ms cpu, 47->47->0 MB, 48 MB goal, 8 P
//gc 9 @0.038s 0%: 0.001+0.060+0.002 ms clock, 0.014+0.025/0.033/0.054+0.022 ms cpu, 47->47->0 MB, 48 MB goal, 8 P
//gc 10 @0.041s 0%: 0.003+0.076+0.003 ms clock, 0.026+0/0.068/0.032+0.026 ms cpu, 47->47->0 MB, 48 MB goal, 8 P
//mem.Alloc: 115568
//mem.TotalAlloc: 500156960
//mem.HeapAlloc 115568
//mem.NumGC: 10
//-----

