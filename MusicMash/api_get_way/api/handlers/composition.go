package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateComposition(gn *gin.Context) {
	var composition pb.CreateCompositionRequest
	err := gn.ShouldBindJSON(&composition)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	response, err := h.Composition.CreateComposition(gn, &composition)
	if err != nil {
		fmt.Println("---------", &response)
		fmt.Println("++++++", err)
		InternalServerError(gn, err)
		return
	}
	Created(gn, err)
	fmt.Println(response)
}

//	func (h *Handler) GetCompositionById(gn *gin.Context) {
//		var id *pb.IdRequest
//		id = gn.Param("id")
//		h.Composition.GetCompositionById(id)
//	}
func (h *Handler) UpdateComposition(gn *gin.Context) {
	var composition *pb.UpdateCompositionRequest
	err := gn.ShouldBindJSON(&composition)
	if err != nil {
		BadRequest(gn, err)
	}
	response, err := h.Composition.UpdateComposition(gn, composition)
	if err != nil {
		InternalServerError(gn, err)
	}
	OK(gn, err)
	fmt.Println(response)

}
func (h *Handler) DeleteComposition(gn *gin.Context) {
	var composition *pb.IdRequest
	err := gn.ShouldBindJSON(&composition)
	if err != nil {
		BadRequest(gn, err)
	}
	response, err := h.Composition.DeleteComposition(gn, composition)
	if err != nil {
		InternalServerError(gn, err)
	}
	OK(gn, err)
	fmt.Println(response)

}
func (h *Handler) GetCompositionByUserId(gn *gin.Context) {
	var composition *pb.IdRequest
	err := gn.ShouldBindJSON(&composition)
	if err != nil {
		BadRequest(gn, err)
	}
	response, err := h.Composition.GetCompositionByUserid(gn, composition)
	if err != nil {
		InternalServerError(gn, err)
	}
	gn.JSON(200, response)

}
