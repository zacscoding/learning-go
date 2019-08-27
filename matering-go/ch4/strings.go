package main

import "fmt"

func main() {
	// \xAB 형태의 기호는 각각 sLiteral을 구성하는 문자 하나 표현
	const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	fmt.Println(sLiteral)
	fmt.Printf("x:%x\n", sLiteral)
	fmt.Printf("sLiteral length:%d\n", len(sLiteral))

	//�B2UP5#P)�
	//x:9942325550352350299c
	//sLiteral length:10

	for i := 0; i < len(sLiteral); i++ {
		fmt.Printf("%x", sLiteral[i])
	}
	fmt.Println()
	//9942325550352350299c

	fmt.Printf("q: %q\n", sLiteral)
	fmt.Printf("+q: %+q\n", sLiteral) // ascii
	fmt.Printf(" x: % x\n", sLiteral) // 바이트 사이의 공백

	fmt.Printf("s: As a string: %s\n", sLiteral)
	//q: "\x99B2UP5#P)\x9c"
	//+q: "\x99B2UP5#P)\x9c"
	//x: 99 42 32 55 50 35 23 50 29 9c
	//s: As a string: �B2UP5#P)�

	s2 := "€£³"
	for x, y := range s2 {
		// U+0058 포맷으로 출력
		fmt.Printf("%#U starts at byte position %d\n", y, x)
	}

	fmt.Printf("s2 length : %d\n", len(s2))
	//U+20AC '€' starts at byte position 0
	//U+00A3 '£' starts at byte position 3
	//U+00B3 '³' starts at byte position 5
	//s2 length : 7

	const s3 = "ab12AB"
	fmt.Println("s3:", s3)
	fmt.Printf("x: % x\n", s3)

	fmt.Printf("s3 length: %d\n", len(s3))

	for i := 0; i < len(s3); i++ {
		fmt.Printf("%x ", s3[i])
	}
	//s3: ab12AB
	//x: 61 62 31 32 41 42
	//s3 length: 6
	//61 62 31 32 41 42

}
