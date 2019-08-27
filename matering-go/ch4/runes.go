package main

import "fmt"

func main() {

	const r1 = '€'
	fmt.Println("(int32) r1:", r1)
	fmt.Printf("(HEX) r1: %x\n", r1)
	fmt.Printf("(as a String) r1: %s\n", r1)
	fmt.Printf("(as a character) r1: %c\n", r1)

	//(int32) r1: 8364
	//(HEX) r1: 20ac
	//(as a String) r1: %!s(int32=8364)
	//(as a character) r1: €

	fmt.Println("A string is a collection of runes:", []byte("Mihalis"))
	aString := []byte("Mihalis")
	for x, y := range aString {
		fmt.Println(x, y)
		fmt.Printf("Char:%c\n", aString[x])
	}
	fmt.Printf("%s\n", aString)
	//A string is a collection of runes: [77 105 104 97 108 105 115]
	//0 77
	//Char:M
	//1 105
	//Char:i
	//2 104
	//Char:h
	//3 97
	//Char:a
	//4 108
	//Char:l
	//5 105
	//Char:i
	//6 115
	//Char:s
	//Mihalis
}
