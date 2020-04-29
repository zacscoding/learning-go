package main

import (
	"context"
	"fmt"
	higrpc "github.com/zacscoding/learning-go/workspace/rpc/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type helloServiceServer struct {
}

func (h *helloServiceServer) Hello(ctx context.Context, req *higrpc.HelloRequest) (*higrpc.HelloResponse, error) {
	fmt.Println("Called Hello() in server ==> first name :", req.FirstName, ", last name :", req.LastName)
	return &higrpc.HelloResponse{Greeting: "Hello " + req.FirstName}, nil
}

func main() {
	fmt.Println(">>>>>> Start gRpc Server1")

	l, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	higrpc.RegisterHelloServiceServer(s, &helloServiceServer{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve grpc server: %v", err)
	}
}
