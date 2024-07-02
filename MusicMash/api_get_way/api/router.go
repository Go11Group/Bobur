package api

import (
	"api_get_way/api/handlers"
	"api_get_way/genproto"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func RouterAPi(compServ genproto.CompositionServiceClient, coll genproto.CollaborationServiceClient, discServ genproto.DiscoveryServiceClient) *gin.Engine {
	router := gin.Default()
	h := handlers.Handler{
		Composition:   compServ,
		Collaboration: coll,
		Discovery:     discServ,
	}
	router.POST("/api/compositions", h.CreateComposition)
	router.GET("/api/compositions/:id")
	router.DELETE(" /api/compositions/:id", h.DeleteComposition)
	router.GET("/api/users/:id/compositions", h.GetCompositionByUserId)
	router.POST(" /api/compositions/:id/tracks", h.CreateTrack)
	router.GET(" /api/compositions/:id/tracks", h.GetTrack)
	router.PUT(" /api/compositions/:id/tracks/:trackId", h.PutTrack)
	router.DELETE(" /api/compositions/:id/tracks/:trackId", h.DeleteTrack)

	router.POST(" /api/collaborations/invite", h.CreateInvite)
	router.PUT("/api/collaborations/invite/:id/respond", h.UpdateInvite)
	router.GET("/api/compositions/:id/collaborators", h.GetCollaboration)
	router.PUT("/api/compositions/:id/collaborators/:userId", h.UpdateCollaboration)
	router.DELETE(" /api/compositions/:id/collaborators/:userId", h.DeleteCollaboration)
	router.POST("/api/compositions/:id/comments", h.CreateComment)
	router.GET("/api/compositions/:id/comments", h.GetComment)

	router.GET("/api/discover/trending")
	router.GET("/api/discover/recommended")
	router.GET("/api/discover/genres/:genre")
	router.GET("/api/search")
	router.POST("/api/compositions/:id/like")
	router.DELETE("/api/compositions/:id/like")
	return router

}
