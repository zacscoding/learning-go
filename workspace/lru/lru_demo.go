package main

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
)

func main() {
	cache, err := lru.NewARC(3)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 7; i++ {
		fmt.Println("-------------------------")
		fmt.Println("Push:", i)
		cache.Add(i, i)
		display(cache)
	}
}

func display(cache *lru.ARCCache) {
	fmt.Println("Size :", cache.Len())
	for _, key := range cache.Keys() {
		fmt.Print(key, " ")
	}
	fmt.Println()
}

//-------------------------
//Push: 0
//Size : 1
//0
//-------------------------
//Push: 1
//Size : 2
//0 1
//-------------------------
//Push: 2
//Size : 3
//0 1 2
//-------------------------
//Push: 3
//Size : 3
//1 2 3
//-------------------------
//Push: 4
//Size : 3
//2 3 4
//-------------------------
//Push: 5
//Size : 3
//3 4 5
//-------------------------
//Push: 6
//Size : 3
//4 5 6
