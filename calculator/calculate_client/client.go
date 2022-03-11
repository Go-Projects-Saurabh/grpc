package main

import (
	"context"
	"fmt"
	"io"
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
	// doUnary(c)
	doServerStreaming(c)
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

func doServerStreaming(c calculatepb.CalculatorServiceClient) {
	fmt.Println("Starting to server streaming RPC ...")
	req := &calculatepb.PrimeRequest{
		Num: 4520,
	}
	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Greet RPS:%v", err)
	}
	fmt.Printf("Prime factors of %v are:\n", req.Num)
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream %v", err)
		}
		fmt.Println(msg.GetResult())
	}
}
