package service

import (
	"context"
	"model/postgres"
	pb " model/genproto/generator"
)

type library struct {
	pb.UnimplementedLibraryServiceServer
	Library *postgres.LibraryRepo
}

func NewLibraryService(db *postgres.LibraryRepo) *library {
	return &library{Library: db}
}

// client dan ma'lumot keladi va id sini qaytaradi
func (s *library) AddBook(conx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	resp, err := s.Library.ADDBOOK(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *library) SearchBook(conx context.Context, req *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	resp, err := s.Library.SearchBook(req)
	if err != nil {
		return nil, err
	}

func(s )