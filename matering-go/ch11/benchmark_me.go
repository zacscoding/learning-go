package main

import "fmt"

func bfibo1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return bfibo1(n-1) + bfibo1(n-2)
	}
}

func bfibo2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return bfibo2(n-1) + bfibo2(n-2)
}

func bfibo3(n int) int {
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
	return fn[n]
}

func main() {
	fmt.Println(bfibo1(40))
	fmt.Println(bfibo2(40))
	fmt.Println(bfibo3(40))
}
