package service

import (
	"context"
	pb "portfolio_service/genproto"
	"portfolio_service/storage"

	"github.com/google/uuid"
)

type ExperienceService struct {
	stg storage.StorageI
	pb.UnimplementedExperiencesServiceServer
}

func NewExperienceService(stg storage.StorageI) *ExperienceService {
	return &ExperienceService{
		stg: stg,
	}
}

func (s *ExperienceService) CreateExperience(ctx context.Context, in *pb.CreateExperienceRequest) (*pb.Experience, error) {
	id := uuid.NewString()
	in.Id = id
	return s.stg.Experience().CreateExperience(in)
}

func (s *ExperienceService) UpdateExperience(ctx context.Context, in *pb.UpdateExperienceRequest) (*pb.Void, error) {
	return s.stg.Experience().UpdateExperience(in)
}

func (s *ExperienceService) DeleteExperience(ctx context.Context, in *pb.ById) (*pb.Void, error) {
	return s.stg.Experience().DeleteExperience(in)
}

func (s *ExperienceService) GetExperience(ctx context.Context, in *pb.ById) (*pb.Experience, error) {
	return s.stg.Experience().GetExperience(in)
}

func (s *ExperienceService) GetExperiences(ctx context.Context, in *pb.ExperienceFilter) (*pb.Experiences, error) {
	return s.stg.Experience().GetExperiences(in)
}
