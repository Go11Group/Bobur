package service

import (
	"context"
	pb "portfolio_service/genproto"
	"portfolio_service/storage"

	"github.com/google/uuid"
)

type EducationService struct {
	stg storage.StorageI
	pb.UnimplementedEducationServiceServer
}

func NewEducationService(stg storage.StorageI) *EducationService {
	return &EducationService{
		stg: stg,
	}
}

func (s *EducationService) CreateEducation(ctx context.Context, in *pb.CreateEducationRequest) (*pb.Education, error) {
	id := uuid.NewString()
	in.Id = id
	return s.stg.Education().CreateEducation(in)
}


func (s *EducationService) UpdateEducation(ctx context.Context, in *pb.UpdateEducationRequest) (*pb.Void, error) {
	return s.stg.Education().UpdateEducation(in)
}


func (s *EducationService) DeleteEducation(ctx context.Context, in *pb.ById) (*pb.Void, error) {
	return s.stg.Education().DeleteEducation(in)
}


func (s *EducationService) GetEducation(ctx context.Context, in *pb.ById) (*pb.Education, error) {
	return s.stg.Education().GetEducation(in)
}


func (s *EducationService) GetEducations(ctx context.Context, in *pb.EducationFilter) (*pb.Educations, error) {
	return s.stg.Education().GetEducations(in)
}

