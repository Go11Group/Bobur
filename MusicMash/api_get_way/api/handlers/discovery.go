package handlers

import (
	pb "api_get_way/genproto"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) GetTrending(gn *gin.Context) {
	response, err := h.Discovery.GetCompositionTrending(gn, &pb.Void{})
	if err != nil {
		InternalServerError(gn, err)
	}
	gn.JSON(200, response)

}
func (h *Handler) GetRecommended(gn *gin.Context) {
	response, err := h.Discovery.GetCompositionRecommend(gn, &pb.Void{})
	if err != nil {
		InternalServerError(gn, err)
	}
	gn.JSON(200, response)

}
func (h *Handler) GetDiscoveryByGenre(gn *gin.Context) {
	var genre *pb.GetGenre
	genre.Genre = gn.Param("genre")
	response, err := h.Discovery.GetCompositionGenre(gn, genre)
	if err != nil {
		InternalServerError(gn, err)
	}
	gn.JSON(200, response)

}
func (h *Handler) SearchDiscovery(gn *gin.Context) {
	var discovery *pb.GetDiscoveryRequest
	discovery.Genre = gn.Query("genre")
	limit, err := strconv.Atoi(gn.Query("limit"))
	if err != nil {
		BadRequest(gn, err)
	}
	discovery.LimitOffset.Limit = int32(limit)
	offset, err := strconv.Atoi(gn.Query("offset"))
	if err != nil {
		BadRequest(gn, err)
	}
	likeCount, err := strconv.Atoi(gn.Query("like_count"))
	if err != nil {
		BadRequest(gn, err)
	}
	discovery.LikeCount = int64(likeCount)
	listenCount, err := strconv.Atoi(gn.Query("listen_count"))
	if err != nil {
		BadRequest(gn, err)
	}
	discovery.ListenCount = int64(listenCount)
	discovery.LimitOffset.Offset = int32(offset)
	response, err := h.Discovery.GetDiscovery(gn, discovery)
	if err != nil {
		InternalServerError(gn, err)
	}
	gn.JSON(200, response)

}
