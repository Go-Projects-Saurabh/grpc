package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/saurabhsingh1408/grpc_greet/calculator/calculatepb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calculate(ctx context.Context, req *calculatepb.SumRequest) (*calculatepb.SumResponse, error) {
	firstnum := req.Sum.Firstnum
	secondnum := req.Sum.Secondnum
	res := &calculatepb.SumResponse{
		Result: firstnum + secondnum,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello! I am a server..")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	calculatepb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
