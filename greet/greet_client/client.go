package main

import (
	"fmt"
	"log"

	"github.com/chanchai-maqe/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

const endpoint = "localhost:50051"

func main() {
	fmt.Print("This message is sent from client\n")
	/*
		Uses WithInsecure() b/c it is the testing and there is no SSL
		required at this moment.
	*/
	cc, err := grpc.Dial(endpoint, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	/*cc will close after greetpb.NewGreetServiceClient*/
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	fmt.Printf("Created client %f :", c)
	fmt.Println()
}
