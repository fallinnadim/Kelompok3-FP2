package repository

import (
	request "fp2/data/request/social_media"
	response "fp2/data/response/social_media"
	"fp2/models"
)

type SocialMediaRepository interface {
	FindAll() []response.AllSocialMediaResponse
	Create(sm request.CreateSocialMediaRequest) models.SocialMedia
	Update(sm request.UpdateSocialMediaRequest) models.SocialMedia
	Delete(id int)
	FindById(id int) (models.SocialMedia, error)
}
