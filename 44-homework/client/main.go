package main

import (
	"context"
	"fmt"
	"log"
	pb "model/genproto/generator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// var (
// 	port = flag.Int("port", 50051, "The server port")

// )

type Server struct {
	pb.UnimplementedGeneratorServer
}

func main() {
	fmt.Println("Client running...")
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGeneratorClient(conn)
	req := &pb.Request{Name: "olma"}

	resp, err := client.RandomPicker(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling RandomPicker: %v", err)
	}

	fmt.Printf("Response: %s\n", resp.Fname)
}
