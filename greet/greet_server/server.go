package main

import (
	"fmt"
	"grpc-demos/greet/greetpb/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreeterServer
}

func main() {
	fmt.Println("Server running")
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Unable to open port 500051", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreeterService(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
