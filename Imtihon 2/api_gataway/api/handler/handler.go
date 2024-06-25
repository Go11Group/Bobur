package handler

import (
	"api_gateway/genproto"

	"google.golang.org/grpc"
)

type Handler struct {
	Education  genproto.EducationServiceClient
	Project    genproto.ProjectsServiceClient
	Skill      genproto.SkillsServiceClient
	Experience genproto.ExperiencesServiceClient
}

func NewHandlerStruct(cl *grpc.ClientConn) *Handler {
	return &Handler{
		Education:  genproto.NewEducationServiceClient(cl),
		Project:    genproto.NewProjectsServiceClient(cl),
		Skill:      genproto.NewSkillsServiceClient(cl),
		Experience: genproto.NewExperiencesServiceClient(cl),
	}
}
