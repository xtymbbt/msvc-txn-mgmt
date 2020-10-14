package main

import (
	pb "../proto"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// server.go

const (
	port = ":1996"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.Person) (*pb.Greeting, error) {
	fmt.Printf("server recieved message: %#v\n\n", in)
	return &pb.Greeting{Message: "Hello " + in.FirstName + " " + in.LastName + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloWorldServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to start server.")
	}
	//s.Serve(lis)
}