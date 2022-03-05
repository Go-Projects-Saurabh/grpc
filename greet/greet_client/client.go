package main

import (
	"context"
	"fmt"
	"log"

	"github.com/saurabhsingh1408/grpc_greet/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello! I am a client.")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Unary RPC ...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Saurabh",
			LastName:  "Singh",
		},
	}

	resp, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPS:%v", err)
	}
	log.Printf("Response from Greet:%v", resp)
}
