package main

import (
	"context"
	"fmt"
	"log"
	pb "model/genproto/generator"
	"model/postgres"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLibraryServiceServer
	Library *postgres.LibraryRepo
}

// client dan ma'lumot keladi va id sini qaytaradi
func (s *server) AddBook(conx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	resp, err := s.Library.ADDBOOK(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// 
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
		panic(err)
	}

	defer db.Close()

	fmt.Println("Server running...")
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	lr := postgres.NewLibrayRepo(db)

	grpcServer := grpc.NewServer()
	pb.RegisterLibraryServiceServer(grpcServer, &server{
		Library: lr,
	})

	fmt.Println("Server is ready to accept requests on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
