package service

import (
	"context"
	pb "portfolio_service/genproto"
	"portfolio_service/storage"

	"github.com/google/uuid"
)

type ProjectService struct {
	stg storage.StorageI
	pb.UnimplementedProjectsServiceServer
}

func NewProjectService(stg storage.StorageI) *ProjectService {
	return &ProjectService{
		stg: stg,
	}
}

func (s *ProjectService) CreateProject(ctx context.Context, in *pb.CreateProjectRequest) (*pb.Void, error) {
	id := uuid.NewString()
	in.Id = id

	_, err := s.stg.Projects().CreateProject(in)

	return &pb.Void{}, err
}

func (s *ProjectService) UpdateProject(ctx context.Context, in *pb.UpdateProjectRequest) (*pb.Void, error) {
	return s.stg.Projects().UpdateProject(in)
}

func (s *ProjectService) DeleteProject(ctx context.Context, in *pb.ById) (*pb.Void, error) {
	return s.stg.Projects().DeleteProject(in)
}

func (s *ProjectService) GetProject(ctx context.Context, in *pb.ById) (*pb.Project, error) {
	return s.stg.Projects().GetProject(in)
}

func (s *ProjectService) GetProjects(ctx context.Context, in *pb.ProjectFilter) (*pb.Projects, error) {
	return s.stg.Projects().GetProjects(in)
}
