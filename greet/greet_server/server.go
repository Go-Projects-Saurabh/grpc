package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/saurabhsingh1408/grpc_greet/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstname := req.Greeting.FirstName
	result := "Hello " + firstname
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*server) GreatManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreatManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)
	firstname := req.Greeting.FirstName
	for i := 0; i < 10; i++ {
		result := "Hello " + firstname + " " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {
	fmt.Println("Hello")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
