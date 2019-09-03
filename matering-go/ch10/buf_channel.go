package main

import "fmt"

func main() {
	// 최대 5개의 정수를 저장
	numbers := make(chan int, 5)
	counter := 10

	for i := 0; i < counter; i++ {
		select {
		case numbers <- i:
		default:
			fmt.Println("Not enough space for", i)
		}
	}

	for i := 0; i < counter+5; i++ {
		select {
		case num := <-numbers:
			fmt.Println(num)
		default:
			fmt.Println("Nothing more to be done!")
			break
		}
	}

	//Not enough space for 5
	//Not enough space for 6
	//Not enough space for 7
	//Not enough space for 8
	//Not enough space for 9
	//0
	//1
	//2
	//3
	//4
	//Nothing more to be done!
	//Nothing more to be done!
	//Nothing more to be done!
	//Nothing more to be done!
	//Nothing more to be done!
	//Nothing more to be done!
	//Nothing more to be done!
	//Nothing more to be done!
	//Nothing more to be done!
	//Nothing more to be done!
}
