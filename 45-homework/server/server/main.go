package main

import (
	"fmt"
	"log"
	"model/postgres"
	"model/service"
	"net"
	pb "model/genproto/generator"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Server running...")
	listener, err := net.Listen("tcp", ":50051") // tarmoq aloqasini o`rnatish uchun kerak
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return
	}

	lib := service.NewLibraryService(postgres.NewLibrayRepo(db))

	s := grpc.NewServer() //yangi grpc server yaratildi

	pb.RegisterLibraryServiceServer(s, lib)

	// bu yerda grpc server ishga tushadi
	if err = s.Serve(listener); err != nil {
		panic(err)
	}

}
