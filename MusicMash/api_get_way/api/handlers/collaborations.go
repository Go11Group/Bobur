package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) GetCollaboration(gn *gin.Context) {
	var collaboration *pb.GetCollaboratorsRequest
	collaboration.UserId = gn.Query("user_id")
	limit, err := strconv.Atoi(gn.Query("limit"))
	if err != nil {
		BadRequest(gn, err)
	}
	collaboration.LimitOffset.Limit = int32(limit)
	offset, err := strconv.Atoi(gn.Query("offset"))
	if err != nil {
		BadRequest(gn, err)
	}
	collaboration.LimitOffset.Offset = int32(offset)
	collaboration.Role = gn.Query("role")
	collaboration.UserId = gn.Query("user_id")
	collaboration.CompositionId = gn.Param("composition_id")
	response, err := h.Collaboration.GetCollaborators(gn, collaboration)
	if err != nil {
		InternalServerError(gn, err)
	}
	gn.JSON(200, response)

}
func (h *Handler) UpdateCollaboration(gn *gin.Context) {
	var collaboration *pb.UpdateCollaborationRequest
	err := gn.ShouldBindJSON(&collaboration)
	if err != nil {
		BadRequest(gn, err)
	}
	collaboration.CompositionId = gn.Param("composition_id")
	collaboration.Userid = gn.Param("user_id")
	response, err := h.Collaboration.UpdateCollaborators(gn, collaboration)
	if err != nil {
		InternalServerError(gn, err)
	}
	fmt.Println(response)
	OK(gn, err)

}
func (h *Handler) DeleteCollaboration(gn *gin.Context) {
	var collaboration *pb.DeleteCollaborationRequest
	collaboration.CompositionId = gn.Param("composition_id")
	collaboration.Userid = gn.Param("user_id")
	response, err := h.Collaboration.DeleteCollaborators(gn, collaboration)
	if err != nil {
		InternalServerError(gn, err)
	}
	fmt.Println(response)
	OK(gn, err)

}
func (h *Handler) CreateComment(gn *gin.Context) {
	var commit *pb.CreateCommitRequest
	err := gn.ShouldBindJSON(&commit)
	if err != nil {
		BadRequest(gn, err)
	}
	commit.CompositionId = gn.Param("composition_id")
	response, err := h.Collaboration.CreateComment(gn, commit)
	if err != nil {
		InternalServerError(gn, err)
	}
	fmt.Println(response)
	OK(gn, err)

}
func (h *Handler) GetComment(gn *gin.Context) {
	var commit *pb.GetCommitRequest
	commit.UserId = gn.Query("user_id")
	limit, err := strconv.Atoi(gn.Query("limit"))
	if err != nil {
		BadRequest(gn, err)
	}
	commit.LimitOffset.Limit = int32(limit)
	offset, err := strconv.Atoi(gn.Query("offset"))
	if err != nil {
		BadRequest(gn, err)
	}
	commit.LimitOffset.Offset = int32(offset)
	commit.CompositionId = gn.Param("composition_id")
	response, err := h.Collaboration.GetComment(gn, commit)
	if err != nil {
		InternalServerError(gn, err)
	}
	gn.JSON(200, response)

}
