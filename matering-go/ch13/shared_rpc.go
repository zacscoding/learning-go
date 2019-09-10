package main

type MyFloats struct {
	A1, A2 float64
}

type MyInterface interface {
	Multiply(arguments *MyFloats, replay *float64) error
	Power(arguments *MyFloats, reply *float64) error
}

//$ mkdir -p ~/go/src/sharedRPC
//$ cp ./shared_rpc.go ~/go/src/sharedRPC/
//$ go install sharedRPC
