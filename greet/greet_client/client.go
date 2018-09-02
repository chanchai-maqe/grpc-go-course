package main

import (
	"context"
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

	/*
		NewGreetServiceClient is a function that return
		GreetServiceClient
	*/
	c := greetpb.NewGreetServiceClient(cc)

	//fmt.Printf("Created client %v :", c)
	fmt.Println()

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {

	fmt.Println("Starting to do a Unary RPC ...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Jack",
			LastName:  "Ryan",
			TagId:     "1100nhh9zxcmasb",
		},
	}

	// Greet function return response and err
	// context comefrom "golang.org/x/net/context"
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Greet function from client: %v", err)
	}

	log.Printf("Response : %v", res.Result)
	log.Println("SHOW RESULT:...")
	//log.Printf("firstname: %v\nlastname: %v\ntagid: %v", res.Result.FirstName, res.Result.LastName, res.Result.TagId)
}
