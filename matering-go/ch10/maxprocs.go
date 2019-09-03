package main

import (
	"fmt"
	"runtime"
)

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

// export GOMAXPROCS=800; go run maxprocs.go
// GOMAXPROCS: 800
func main() {
	fmt.Printf("GOMAXPROCS: %d\n", getGOMAXPROCS())
}
