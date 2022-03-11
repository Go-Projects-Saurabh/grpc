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

func (*server) PrimeNumberDecomposition(req *calculatepb.PrimeRequest, stream calculatepb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("PrimeNumberDecomposition function was invoked with %v\n", req)
	var k int64
	k = 2
	N := req.Num
	for N > 1 {
		if N%k == 0 {
			res := &calculatepb.PrimeResponse{
				Result: int64(k),
			}
			stream.Send(res)
			N = N / k
		} else {
			k = k + 1
		}
	}
	return nil
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
