package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/chanchai-maqe/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

// Implement GreetServiceServer to *server
// by add Greet() method to server struct{}
/*
	It means server recieved GreetRequest from client
	and it has to return GreetResponse as &

	Note: if function return type

*/
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	firstname := req.GetGreeting().GetFirstName()
	lastname := req.GetGreeting().GetLastName()
	tagid := req.GetGreeting().GetTagId()

	result := firstname + "\t" + lastname + "\t" + tagid

	/*
		GreetResponse is a struct that has "Result" string inside
	*/
	// res := &greetpb.GreetResponse{
	// 	Result: result,
	// }

	return &greetpb.GreetResponse{
		Result: result,
	}, nil

}

func main() {
	fmt.Println("Server is running")

	// NOTE: 50051 is default port for grpc
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// create grpc server variable as NewServer()
	s := grpc.NewServer()

	// register server to greetpb variable
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
