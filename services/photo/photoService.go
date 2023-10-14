package services

import (
	request "fp2/data/request/photo"
	response "fp2/data/response/photo"
)

type PhotoService interface {
	Post(p request.CreatePhotoRequest) (response.CreatedPhotoResponse, error)
	GetAll() []response.AllPhotoResponse
	Update(p request.UpdatePhotoRequest) (response.UpdatedPhotoResponse, error)
	Delete(id int) error
}
