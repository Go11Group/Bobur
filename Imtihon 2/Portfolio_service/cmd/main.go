package main

import (
	"log"
	"net"
	pb "portfolio_service/genproto"
	"portfolio_service/service"
	"portfolio_service/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	liss, err := net.Listen("tcp", ":8088")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterEducationServiceServer(s, service.NewEducationService(db))
	pb.RegisterExperiencesServiceServer(s, service.NewExperienceService(db))
	pb.RegisterProjectsServiceServer(s, service.NewProjectService(db))
	pb.RegisterSkillsServiceServer(s, service.NewSkillService(db))

	log.Printf("server listening at %v", liss.Addr())

	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
