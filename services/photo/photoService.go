package services

import "fp2/dto"

type PhotoService interface {
	Post(p dto.CreatePhotoRequest) (dto.CreatedPhotoResponse, error)
	GetAll() []dto.AllPhotoResponse
	Update(p dto.UpdatePhotoRequest) (dto.UpdatedPhotoResponse, error)
	Delete(id int) error
}
