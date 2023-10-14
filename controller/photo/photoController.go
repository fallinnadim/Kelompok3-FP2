package controller

import (
	request "fp2/data/request/photo"
	response "fp2/data/response/users"
	"fp2/helper"
	services "fp2/services/photo"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	PhotoService services.PhotoService
}

func NewPhotoController(s services.PhotoService) *PhotoController {
	return &PhotoController{PhotoService: s}
}

func (p *PhotoController) CreatePhoto(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Invalid token",
		})
		return
	}
	// panggil service
	createPhoto := request.CreatePhotoRequest{}
	err := ctx.ShouldBindJSON(&createPhoto)
	if err != nil {
		webResponse := response.FailedResponse{
			Status:  false,
			Message: helper.ParseError(err),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	createPhoto.User_Id = userId.(int)
	result, errCreate := p.PhotoService.Post(createPhoto)
	// return response
	if errCreate != nil {
		webResponse := response.FailedResponse{
			Status:  false,
			Message: helper.ParseError(errCreate),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

func (p *PhotoController) GetAllPhoto(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get all",
	})
}

func (p *PhotoController) UpdatePhoto(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update",
	})
}

func (p *PhotoController) DeletePhoto(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete",
	})
}
