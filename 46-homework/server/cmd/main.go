package main

import (
	"fmt"
	"log"
	"net"
	"project/postgres"
	"project/service"
	pb "project/weather/generator"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("error --> ", err)
		return
	}
	defer db.Close()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("error --> ", err)
		return
	}
	fmt.Println("Server lesterner 500051...")
	lib := service.NewWeatherSerice(postgres.NewWeatherRepo(db))

	s := grpc.NewServer() // yangi grpc server yaratildi

	pb.RegisterWeatherServiceServer(s, lib) // serverda service ni metodlari ishlash uchuin register qilinadi

	if err = s.Serve(listener); err != nil {
		panic(err)
	}

}


