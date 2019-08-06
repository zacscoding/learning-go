package main

import "fmt"

func main() {
	v1 := "123"
	v2 := 123
	v3 := "Have a nice day\n"
	v4 := "abc"

	fmt.Print(v1, v2, v3, v4)
	fmt.Println()
	fmt.Println(v1, v2, v3, v4)
	fmt.Print(v1, " ", v2, v3, " ", v4, "\n")
	fmt.Printf("%s%d %s %s\n", v1, v2, v3, v4)

	// Output
	//123123Have a nice day
	//abc
	//123 123 Have a nice day
	//abc
	//123 123Have a nice day
	//abc
	//123123 Have a nice day
	//abc

}
