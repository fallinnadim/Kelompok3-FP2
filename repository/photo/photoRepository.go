package repository

import (
	request "fp2/data/request/photo"
	response "fp2/data/response/photo"
	"fp2/models"
)

type PhotoRepository interface {
	FindAll() []response.AllPhotoResponse
	Create(sm request.CreatePhotoRequest) models.SocialMedia
	Update(sm request.UpdatePhotoRequest) models.SocialMedia
	Delete(id int)
	FindById(id int) (models.Photo, error)
}