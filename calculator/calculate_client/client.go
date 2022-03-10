package main

import (
	"context"
	"fmt"
	"log"

	"github.com/saurabhsingh1408/grpc_greet/calculator/calculatepb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello! I am a client ...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect: %v", err)
	}

	defer cc.Close()

	c := calculatepb.NewCalculatorServiceClient(cc)
	doUnary(c)
}

func doUnary(c calculatepb.CalculatorServiceClient) {
	fmt.Println("Starting to do Unary RPC ...")
	req := &calculatepb.SumRequest{
		Sum: &calculatepb.Sum{
			Firstnum:  20,
			Secondnum: 10,
		},
	}

	resp, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC:%v", err)
	}
	log.Printf("Response from Calculate:%v", resp)
}
