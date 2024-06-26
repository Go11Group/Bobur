package main

import (
	"context"
	"fmt"
	"log"
	pb "model/genproto/generator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	pb.UnimplementedLibraryServiceServer
}

func main() {
	fmt.Println("Client running...")
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewLibraryServiceClient(conn)

	// book yaratish
	req := &pb.AddBookRequest{
		Title:         "asdfasd",
		Author:        "bobur",
		YearPublished: 2024,
	}

	resp, err := client.AddBook(context.Background(), req)

	if err != nil {
		log.Fatalf("Error calling AddBook: %v", err)
	}

	fmt.Printf("Response: %s\n", resp)
}
