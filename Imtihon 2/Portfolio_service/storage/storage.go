package storage

import (
	pb "portfolio_service/genproto"
)

type StorageI interface {
	Education() EducationI
	Experience() ExperienceI
	Projects() ProjectsI
	Skills() SkillsI
}

type EducationI interface {
	CreateEducation(education *pb.CreateEducationRequest) (*pb.Education, error)
	UpdateEducation(education *pb.UpdateEducationRequest) (*pb.Void, error)
	DeleteEducation(education *pb.ById) (*pb.Void, error)
	GetEducation(flt *pb.ById) (*pb.Education, error)
	GetEducations(flt *pb.EducationFilter) (*pb.Educations, error)
}

type ExperienceI interface {
	CreateExperience(experience *pb.CreateExperienceRequest) (*pb.Experience, error)
	UpdateExperience(experience *pb.UpdateExperienceRequest) (*pb.Void, error)
	DeleteExperience(experience *pb.ById) (*pb.Void, error)
	GetExperience(flt *pb.ById) (*pb.Experience, error)
	GetExperiences(flt *pb.ExperienceFilter) (*pb.Experiences, error)
}

type ProjectsI interface {
	CreateProject(project *pb.CreateProjectRequest) (*pb.Project, error)
	UpdateProject(project *pb.UpdateProjectRequest) (*pb.Void, error)
	DeleteProject(project *pb.ById) (*pb.Void, error)
	GetProject(flt *pb.ById) (*pb.Project, error)
	GetProjects(flt *pb.ProjectFilter) (*pb.Projects, error)
}

type SkillsI interface {
	CreateSkill(skill *pb.CreateSkillRequest) (*pb.Skill, error)
	UpdateSkill(skill *pb.UpdateSkillRequest) (*pb.Void, error)
	DeleteSkill(skill *pb.ById) (*pb.Void, error)
	GetSkill(flt *pb.ById) (*pb.Skill, error)
	GetSkills(flt *pb.SkillFilter) (*pb.Skills, error)
}