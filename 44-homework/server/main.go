package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "model/genproto/generator"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGeneratorServer
}



func (s *server) RandomPicker(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	var surnames = map[string]string{
		"olma":"appel",
		"son":"integer",
		"gul":"flower",
		"telefon":"phone",
	}
	name := in.Name
	fmt.Println(name, surnames[name], 111)
	return &pb.Response{Fname: surnames[name]}, nil
}

func main() {
	fmt.Println("Server running...")
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGeneratorServer(grpcServer, &server{})

	fmt.Println("Server is ready to accept requests on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
