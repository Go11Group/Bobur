package service

import (
	"context"
	pb "portfolio_service/genproto"
	"portfolio_service/storage"

	"github.com/google/uuid"
)

type SkillService struct {
	stg storage.StorageI
	pb.UnimplementedSkillsServiceServer
}



func NewSkillService(stg storage.StorageI) *SkillService {
	return &SkillService{
		stg: stg,
	}
}


func (s *SkillService) CreateSkill(ctx context.Context, in *pb.CreateSkillRequest) (*pb.Skill, error) {
	id := uuid.NewString()
	in.Id = id
	return s.stg.Skills().CreateSkill(in)
}


func (s *SkillService) UpdateSkill(ctx context.Context, in *pb.UpdateSkillRequest) (*pb.Void, error) {
	return s.stg.Skills().UpdateSkill(in)
}


func (s *SkillService) DeleteSkill(ctx context.Context, in *pb.ById) (*pb.Void, error) {
	return s.stg.Skills().DeleteSkill(in)
}


func (s *SkillService) GetSkill(ctx context.Context, in *pb.ById) (*pb.Skill, error) {
	return s.stg.Skills().GetSkill(in)
}

func (s *SkillService) GetSkills(ctx context.Context, in *pb.SkillFilter) (*pb.Skills, error) {
	return s.stg.Skills().GetSkills(in)
}
