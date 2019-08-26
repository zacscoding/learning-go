package main

import "fmt"

// $ go tool compile -W node_tree.go
// $ go tool compile -W node_tree.go | grep before
// $ go tool compile -W node_tree.go | grep after
func main() {
	fmt.Println("Hello there!")
}
