package services

import (
	request "fp2/data/request/social_media"
	response "fp2/data/response/social_media"
)

type SocialMediaService interface {
	Post(sm request.CreateSocialMediaRequest) (response.CreatedSocialMediaResponse, error)
	GetAll(userId int)
	Update(sm request.UpdateSocialMediaRequest) (response.UpdatedSocialMediaResponse, error)
	Delete(id int)
}
