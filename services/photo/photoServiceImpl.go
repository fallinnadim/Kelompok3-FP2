package services

import (
	"fp2/dto"
	repository "fp2/repository/photo"
	"time"

	"github.com/go-playground/validator/v10"
)

type PhotoServiceImpl struct {
	PhotoRepository repository.PhotoRepository
	Validate        *validator.Validate
}

// Delete implements PhotoService.
func (p *PhotoServiceImpl) Delete(id int) error {
	// Panggil service
	err := p.PhotoRepository.Delete(id)
	return err
}

// GetAll implements PhotoService.
func (p *PhotoServiceImpl) GetAll() []dto.AllPhotoResponse {
	result := p.PhotoRepository.FindAll()
	return result
}

// Post implements PhotoService.
func (p *PhotoServiceImpl) Post(cp dto.CreatePhotoRequest) (dto.CreatedPhotoResponse, error) {
	// Validasi Struct
	errValidation := p.Validate.Struct(cp)
	if errValidation != nil {
		return dto.CreatedPhotoResponse{}, errValidation
	}
	cp.Created_At = time.Now().Format("2006-01-02")
	cp.Updated_At = time.Now().Format("2006-01-02")
	// Panggil Repository
	result := p.PhotoRepository.Create(cp)
	// Return
	resp := dto.CreatedPhotoResponse{
		Id:         result.Id,
		Title:      result.Title,
		Caption:    result.Caption,
		Photo_Url:  result.Photo_Url,
		User_Id:    result.User_Id,
		Created_At: result.Created_At,
	}
	return resp, nil
}

// Update implements PhotoService.
func (p *PhotoServiceImpl) Update(cp dto.UpdatePhotoRequest) (dto.UpdatedPhotoResponse, error) {
	// Validasi Struct
	errValidation := p.Validate.Struct(cp)
	if errValidation != nil {
		return dto.UpdatedPhotoResponse{}, errValidation
	}
	cp.Updated_At = time.Now().Format("2006-01-02")
	// Panggil service
	result := p.PhotoRepository.Update(cp)
	updatePhoto := dto.UpdatedPhotoResponse{
		Id:         result.Id,
		Title:      result.Title,
		Caption:    result.Caption,
		Photo_Url:  result.Photo_Url,
		User_Id:    result.User_Id,
		Updated_At: result.Updated_At,
	}
	return updatePhoto, nil
}

func NewPhotoServiceImpl(p repository.PhotoRepository, v *validator.Validate) PhotoService {
	return &PhotoServiceImpl{
		PhotoRepository: p,
		Validate:        v,
	}
}
