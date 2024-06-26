package main

import (
	"context"
	"fmt"
	"log"
	pb "model/genproto/generator"
	// pb "model/genproto/generator"
	"model/postgres"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLibraryServiceServer

	Library *postgres.LibraryRepo
}

func (s *server) AddBook(conx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	resp, err := s.Library.ADDBOOK(req)
	if err != nil {
		fmt.Println(err, "1-1-1")
		return nil, err
	}
	return resp, nil
}

func (s *server) SearchBook(conx context.Context, req *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	resp, err := s.Library.SearchBook(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *server) BorrowBook(conx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	return nil, nil
}

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer db.Close()

	fmt.Println("Server running...")
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLibraryServiceServer(grpcServer, &server{})

	fmt.Println("Server is ready to accept requests on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
