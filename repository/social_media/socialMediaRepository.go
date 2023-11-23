package repository

import (
	"fp2/dto"
	"fp2/entity"
)

type SocialMediaRepository interface {
	FindAll() []dto.AllSocialMediaResponse
	Create(sm dto.CreateSocialMediaRequest) entity.SocialMedia
	Update(sm dto.UpdateSocialMediaRequest) entity.SocialMedia
	Delete(id int)
	FindById(id int) (entity.SocialMedia, error)
}
