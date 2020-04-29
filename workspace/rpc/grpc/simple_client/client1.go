package main

import (
	"context"
	"fmt"
	higrpc "github.com/zacscoding/learning-go/workspace/rpc/grpc/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := higrpc.NewHelloServiceClient(conn)

	res, err := c.Hello(context.Background(), &higrpc.HelloRequest{
		FirstName: "Zac",
		LastName:  "Codding",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("===> Response of grpc :", res.GetGreeting())

}
