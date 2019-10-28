package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func main() {
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			// producer를 1초동안 대기하게 해서 observer 고루틴보다 덜 활동적이도록 만듬
			time.Sleep(1)
		}
	}
	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	// 1 writer, count : reader 로 시간 체크
	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}
		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}

	// Output:
	//Readers  RWMutex       Mutex
	//1        8.406µs       2.628µs
	//2        3.065µs       2.468µs
	//4        5.518µs       2.962µs
	//8        30.769µs      67.279µs
	//16       59.118µs      25.704µs
	//32       56.141µs      42.772µs
	//64       66.212µs      37.469µs
	//128      121.449µs     58.241µs
	//256      125.213µs     86.835µs
	//512      242.506µs     172.107µs
	//1024     248.849µs     352.371µs
	//2048     495.776µs     2.451669ms
	//4096     871.525µs     1.169334ms
	//8192     1.871631ms    2.217825ms
	//16384    3.749914ms    4.447222ms
	//32768    7.399789ms    7.716905ms
	//65536    14.855402ms   18.386883ms
	//131072   30.018601ms   36.159083ms
	//262144   59.834444ms   74.653077ms
	//524288   119.487085ms  148.24082ms
}
