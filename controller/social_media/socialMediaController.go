package controller

import (
	request "fp2/data/request/social_media"
	response "fp2/data/response/users"
	"fp2/helper"
	services "fp2/services/social_media"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SocialMediaController struct {
	SocialMediaService services.SocialMediaService
}

func NewSocialMediaController(s services.SocialMediaService) *SocialMediaController {
	return &SocialMediaController{SocialMediaService: s}
}

func (s *SocialMediaController) CreateSocialMedia(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Invalid token",
		})
		return
	}
	// panggil service
	createSocialMedia := request.CreateSocialMediaRequest{}
	err := ctx.ShouldBindJSON(&createSocialMedia)
	if err != nil {
		webResponse := response.FailedResponse{
			Status:  false,
			Message: helper.ParseError(err),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	createSocialMedia.User_Id = userId.(int)
	result, errCreate := s.SocialMediaService.Post(createSocialMedia)
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
