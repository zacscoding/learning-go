package main

import (
	"context"
	"fmt"
	higrpc "github.com/zacscoding/learning-go/workspace/rpc/grpc/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
	"sync"
	"time"
)

type runTest func(c higrpc.PersonRouteClient) error

type TestSet struct {
	title string
	test  runTest
}

func main() {
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := higrpc.NewPersonRouteClient(conn)
	tests := []TestSet{
		{
			"GetPerson for simple rpc",
			GetPerson(),
		},
		{
			"ListPerson for server to client streaming RPC",
			ListPerson(),
		},
		{
			"SavePerson for client to server streaming RPC",
			SavePerson(),
		},
		{
			"GetPersonChat for bidirectional streaming RPC",
			GetPersonChat(),
		},
	}

	for _, test := range tests {
		fmt.Println("==========================")
		fmt.Println(test.title)
		fmt.Println("--------------------------")
		fmt.Println("Error :", test.test(c))
	}
}

func GetPerson() runTest {
	return func(c higrpc.PersonRouteClient) error {
		res, err := c.GetPerson(context.Background(), &higrpc.PersonQuery{
			Id:   1,
			Name: "Zaccoding",
		})
		fmt.Println("Receive >", res.String())
		return err
	}
}

func ListPerson() runTest {
	return func(c higrpc.PersonRouteClient) error {
		stream, err := c.ListPerson(context.Background(), &higrpc.PersonQuery{
			Id:   2,
			Name: "Zaccoding",
		})
		if err != nil {
			return err
		}

		for {
			res, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("Complete to receive")
				break
			}
			if err != nil {
				return err
			}
			fmt.Println("Receive >", res.String())
		}
		return nil
	}
}

func SavePerson() runTest {
	return func(c higrpc.PersonRouteClient) error {
		stream, err := c.SavePerson(context.Background())
		if err != nil {
			return err
		}
		for i := 1; i < 5; i++ {
			stream.Send(&higrpc.PersonRequest{
				Name: "Zaccoding" + strconv.Itoa(i),
				Age:  int32(i),
			})
		}
		res, err := stream.CloseAndRecv()
		if err != nil {
			return nil
		}
		fmt.Println("Receive >", res.String())
		return nil
	}
}

func GetPersonChat() runTest {
	return func(c higrpc.PersonRouteClient) error {
		stream, err := c.GetPersonChat(context.Background())
		if err != nil {
			return nil
		}
		var wait sync.WaitGroup
		wait.Add(1)
		go func() {
			defer wait.Done()
			for {
				res, err := stream.Recv()
				if err == io.EOF {
					fmt.Println("Complete to receive")
					return
				}
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("Receive >>", res.String())
			}
		}()

		for i := 0; i < 5; i++ {
			stream.Send(&higrpc.PersonQuery{
				Id:   int64(i),
				Name: "Zaccoding" + strconv.Itoa(i),
			})
			time.Sleep(1 * time.Second)
		}
		wait.Wait()
		return nil
	}
}
