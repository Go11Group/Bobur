package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateInvite(gn *gin.Context) {
	var invite *pb.CreateInviteRequest
	err := gn.ShouldBindJSON(&invite)
	if err != nil {
		BadRequest(gn, err)
	}
	response, err := h.Collaboration.CreateInvite(gn, invite)
	if err != nil {
		InternalServerError(gn, err)
	}
	Created(gn, err)
	fmt.Println(response)

}
func (h *Handler) UpdateInvite(gn *gin.Context) {
	var invite *pb.UpdateInviteRequest
	err := gn.ShouldBindJSON(&invite)
	if err != nil {
		BadRequest(gn, err)
	}
	invite.Id = gn.Param("id")
	response, err := h.Collaboration.UpdateInvite(gn, invite)
	if err != nil {
		InternalServerError(gn, err)
	}
	OK(gn, err)
	fmt.Println(response)

}
