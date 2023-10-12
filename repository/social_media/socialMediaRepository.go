package repository

import (
	request "fp2/data/request/social_media"
	"fp2/models"
)

type SocialMediaRepository interface {
	FindAll(id int) []models.SocialMedia
	Create(sm request.CreateSocialMediaRequest) models.SocialMedia
	Update(sm request.UpdateSocialMediaRequest) models.SocialMedia
	Delete(id int)
}
