package repository

import (
	"fp2/dto"
	"fp2/entity"
)

type PhotoRepository interface {
	FindAll() []dto.AllPhotoResponse
	Create(sm dto.CreatePhotoRequest) entity.Photo
	Update(sm dto.UpdatePhotoRequest) entity.Photo
	Delete(id int) error
	FindById(id int) (entity.Photo, error)
}
