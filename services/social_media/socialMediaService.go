package services

import "fp2/dto"

type SocialMediaService interface {
	Post(sm dto.CreateSocialMediaRequest) (dto.CreatedSocialMediaResponse, error)
	GetAll() []dto.AllSocialMediaResponse
	Update(sm dto.UpdateSocialMediaRequest) (dto.UpdatedSocialMediaResponse, error)
	Delete(smId int) error
}
