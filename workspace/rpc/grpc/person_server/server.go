package main

import (
	"context"
	"fmt"
	higrpc "github.com/zacscoding/learning-go/workspace/rpc/grpc/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type PersonServer struct {
}

// GetPerson is test rpc for simple rpc i.e no stream
func (p *PersonServer) GetPerson(ctx context.Context, query *higrpc.PersonQuery) (*higrpc.PersonResponse, error) {
	fmt.Println("PersonServer::GetPerson is called. id :", query.Id, ", name :", query.Name)

	if rand.Intn(100) < 5 {
		return &higrpc.PersonResponse{}, nil
	}

	name := query.Name
	if name == "" {
		name = "Zaccoding"
	}

	return &higrpc.PersonResponse{
		Id:   query.Id,
		Name: name,
		Age:  15,
	}, nil
}

// ListPerson is test rpc for stream of server side
func (p *PersonServer) ListPerson(query *higrpc.PersonQuery, stream higrpc.PersonRoute_ListPersonServer) error {
	fmt.Println("PersonServer::ListPerson is called. id :", query.Id, ", name :", query.Name)

	// ignore query fields, only send hard coded responses
	for i := 0; i < 5; i++ {
		_ = stream.Send(&higrpc.PersonResponse{
			Id:   int64(i),
			Name: "Person" + strconv.Itoa(i),
			Age:  int32(i),
		})
	}
	return nil
}

// SavePerson is test rpc for stream of client side
func (p *PersonServer) SavePerson(stream higrpc.PersonRoute_SavePersonServer) error {
	fmt.Println("PersonServer::SavePerson is called.")

	now := time.Now()
	count := int32(0)
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		count++
		fmt.Println(">> Name :", in.Name, ", Age :", in.Age)
	}
	e := now.Sub(time.Now())
	return stream.SendAndClose(&higrpc.PersonSaveSummary{
		Trial:   count,
		Success: count,
		Fail:    0,
		Elapsed: int64(e),
	})
}

// GetPersonChat is test rpc for bidirectional stream
func (p *PersonServer) GetPersonChat(stream higrpc.PersonRoute_GetPersonChatServer) error {
	fmt.Println("PersonServer::GetPersonChat is called.")

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(">> Id :", in.Id, ", Name:", in.Name)
		_ = stream.Send(&higrpc.PersonResponse{
			Id:   in.Id,
			Name: in.Name,
			Age:  int32(in.Id),
		})
	}
	fmt.Println(">>> Complete")
	return nil
}

func main() {
	fmt.Println(">>>>>> Start gRpc PersonServer")

	l, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	higrpc.RegisterPersonRouteServer(s, &PersonServer{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve grpc server: %v", err)
	}
}
