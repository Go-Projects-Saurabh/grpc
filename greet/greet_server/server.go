package main

import (
	"fmt"
	"log"
	"net"

	"github.com/saurabhsingh1408/grpc_greet/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Hello")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println(lis)
	grpc.NewServer()
	greetpb.RegisterGreetServiceServer()

}
