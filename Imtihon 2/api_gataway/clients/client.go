package clients

import (
	pb "api_gateway/genproto"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientI struct {
	EducationClient  pb.EducationServiceClient
	ExperienceClient pb.ExperiencesServiceClient
	ProjectClient    pb.ProjectsServiceClient
	SkillClient      pb.SkillsServiceClient
}

func NewClient() *ClientI {
	conn, err := grpc.NewClient("localhost:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		slog.Error("Error while dialing", err)
	}

	eduC := pb.NewEducationServiceClient(conn)
	exC := pb.NewExperiencesServiceClient(conn)
	prC := pb.NewProjectsServiceClient(conn)
	skC := pb.NewSkillsServiceClient(conn)

	return &ClientI{
		EducationClient:  eduC,
		ExperienceClient: exC,
		ProjectClient:    prC,
		SkillClient:      skC,
	}

}
