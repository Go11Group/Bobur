package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) CreateTrack(gn *gin.Context) {
	var track *pb.CreateTrackRequest
	err := gn.ShouldBindJSON(&track)
	if err != nil {
		BadRequest(gn, err)
	}
	track.CompositionId = gn.Param("id")
	response, err := h.Composition.CreateTrack(gn, track)
	if err != nil {
		InternalServerError(gn, err)
	}
	Created(gn, err)
	fmt.Println(response)

}
func (h *Handler) GetTrack(gn *gin.Context) {
	var track *pb.GetTrackRequest
	track.Userid = gn.Query("user_id")
	limit, err := strconv.Atoi(gn.Query("limit"))
	if err != nil {
		BadRequest(gn, err)
	}
	track.LimitOffset.Limit = int32(limit)
	offset, err := strconv.Atoi(gn.Query("offset"))
	if err != nil {
		BadRequest(gn, err)
	}
	track.LimitOffset.Offset = int32(offset)
	track.CompositionId = gn.Param("id")
	track.FileUrl = gn.Param("file_url")
	track.Title = gn.Query("title")
	track.CompositionId = gn.Param("composition_id")
	response, err := h.Composition.GetTrack(gn, track)
	if err != nil {
		InternalServerError(gn, err)
	}

	gn.JSON(200, response)

}

func (h *Handler) PutTrack(gn *gin.Context) {
	var track *pb.UpdateTrackRequest
	track.CompositionId = gn.Param("composition_id")
	track.Id = gn.Param("track_id")
	response, err := h.Composition.UpdateTrack(gn, track)
	if err != nil {
		InternalServerError(gn, err)
	}
	fmt.Println(response)
	OK(gn, err)

}
func (h *Handler) DeleteTrack(gn *gin.Context) {
	var track *pb.DeleteTrackRequest
	track.CompositionId = gn.Param("composition_id")
	track.TrackId = gn.Param("track_id")
	response, err := h.Composition.DeleteTrack(gn, track)
	if err != nil {
		InternalServerError(gn, err)
	}
	fmt.Println(response)
	OK(gn, err)

}
