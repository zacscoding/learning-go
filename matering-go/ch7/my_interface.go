package main

type Shape interface {
	Area() float64
	Perimeter() float64
}

//$ mkdir ~/go/src/my_interface
//$ cp ./my_interface.go ~/go/src/my_interface/
//$ go install my_interface
