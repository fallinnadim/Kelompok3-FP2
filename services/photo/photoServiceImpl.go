package services

import (
	request "fp2/data/request/photo"
	response "fp2/data/response/photo"
	repository "fp2/repository/photo"

	"github.com/go-playground/validator/v10"
)

type PhotoServiceImpl struct {
	PhotoRepository repository.PhotoRepository
	Validate        *validator.Validate
}

// Delete implements PhotoService.
func (*PhotoServiceImpl) Delete(id int) error {
	panic("unimplemented")
}

// GetAll implements PhotoService.
func (*PhotoServiceImpl) GetAll() []response.AllPhotoResponse {
	panic("unimplemented")
}

// Post implements PhotoService.
func (*PhotoServiceImpl) Post(p request.CreatePhotoRequest) (response.CreatedPhotoResponse, error) {
	panic("unimplemented")
}

// Update implements PhotoService.
func (*PhotoServiceImpl) Update(p request.UpdatePhotoRequest) (response.UpdatedPhotoResponse, error) {
	panic("unimplemented")
}

func NewPhotoServiceImpl(p repository.PhotoRepository, v *validator.Validate) PhotoService {
	return &PhotoServiceImpl{
		PhotoRepository: p,
		Validate:        v,
	}
}
