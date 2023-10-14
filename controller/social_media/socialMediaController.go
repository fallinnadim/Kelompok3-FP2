package controller

import (
	request "fp2/data/request/social_media"
	response "fp2/data/response/users"
	"fp2/helper"
	services "fp2/services/social_media"
	"net/http"
	"strconv"

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

func (s *SocialMediaController) GetAllSocialMedia(ctx *gin.Context) {
	webResponse := s.SocialMediaService.GetAll()
	ctx.JSON(http.StatusOK, gin.H{
		"social_medias": webResponse,
	})
}

func (s *SocialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	smId, _ := strconv.Atoi(ctx.Param("socialMediaId"))
	// panggil service
	updateSocialMediaRequest := request.UpdateSocialMediaRequest{}
	err := ctx.ShouldBindJSON(&updateSocialMediaRequest)
	if err != nil {
		webResponse := response.FailedResponse{
			Status:  false,
			Message: helper.ParseError(err),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	updateSocialMediaRequest.Id = smId
	updateSocialMediaRequest.User_Id = userId.(int)
	result, errUpdate := s.SocialMediaService.Update(updateSocialMediaRequest)
	// return response
	if errUpdate != nil {
		webResponse := response.FailedResponse{
			Status:  false,
			Message: helper.ParseError(errUpdate),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (s *SocialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	smId, _ := strconv.Atoi(ctx.Param("socialMediaId"))
	// panggil service
	errDelete := s.SocialMediaService.Delete(smId)
	// return response
	if errDelete != nil {
		webResponse := response.FailedResponse{
			Status:  false,
			Message: helper.ParseError(errDelete),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
