package main

import (
	"context"
	"fmt"
	"io"
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

	// doUnary(c)
	doServerStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Unary RPC ...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{FirstName: "Saurabh", LastName: "Singh"},
	}

	resp, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPS:%v", err)
	}
	log.Printf("Response from Greet:%v", resp)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to server streaming RPC ...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{FirstName: "Saurabh", LastName: "Singh"},
	}
	resStream, err := c.GreatManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Greet RPS:%v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream %v", err)
		}
		log.Printf("Response from GreetManyTimes:%v", msg.GetResult())
	}
}
