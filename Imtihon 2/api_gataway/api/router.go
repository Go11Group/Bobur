package api

import (
	"api_gateway/api/handler"
	_ "api_gateway/docs"

	"github.com/gin-contrib/cors"
	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title Portfolio swagger UI
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func InitRouter(p *grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	h := handler.NewHandlerStruct(p)
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}
	router.Use(cors.New(corsConfig))

	education := router.Group("/education")
	{
		education.POST("/create", h.CreateEducation)
		education.PUT("/update", h.UpdateEducation)
		education.DELETE("/delete/:id", h.DeleteEducation)
		education.GET("/getallbyuser", h.GetEducationsByUser)
		education.GET("/get/:id", h.GetEducation)
		education.GET("/getall", h.GetEducations)
	}

	skill := router.Group("/skill")
	{
		skill.POST("/create", h.CreateSkill)
		skill.PUT("/update", h.UpdateSkill)
		skill.DELETE("/delete/:id", h.DeleteSkill)
		skill.GET("/getallbyuser", h.GetSkillsByUser)
		skill.GET("/get/:id", h.GetSkill)
		skill.GET("/getall", h.GetSkills)
	}

	experience := router.Group("/experience")
	{
		experience.POST("/create", h.CreateExperience)
		experience.PUT("/update", h.UpdateExperience)
		experience.DELETE("/delete/:id", h.DeleteExperience)
		experience.GET("/getallbyuser", h.GetExperiencesByUser)
		experience.GET("/get/:id", h.GetExperience)
		experience.GET("/getall", h.GetExperiences)
	}

	project := router.Group("/project")
	{
		project.POST("/create", h.CreateProject)
		project.PUT("/update", h.UpdateProject)
		project.DELETE("/delete/:id", h.DeleteProject)
		project.GET("/getallbyuser", h.GetProjectsByUser)
		project.GET("/get/:id", h.GetProject)
		project.GET("/getall", h.GetProjects)
	}

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router

}
