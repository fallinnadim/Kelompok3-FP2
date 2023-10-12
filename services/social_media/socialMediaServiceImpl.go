package services

import (
	request "fp2/data/request/social_media"
	response "fp2/data/response/social_media"
	repository "fp2/repository/social_media"
	"time"

	"github.com/go-playground/validator/v10"
)

type SocialMediaServiceImpl struct {
	SocialMediaRepository repository.SocialMediaRepository
	Validate              *validator.Validate
}

// Delete implements SocialMediaService.
func (*SocialMediaServiceImpl) Delete(id int) {
	panic("unimplemented")
}

// GetAll implements SocialMediaService.
func (*SocialMediaServiceImpl) GetAll(userId int) {
	panic("unimplemented")
}

// Post implements SocialMediaService.
func (s *SocialMediaServiceImpl) Post(sm request.CreateSocialMediaRequest) (response.CreatedSocialMediaResponse, error) {
	// Validasi Struct
	errValidation := s.Validate.Struct(sm)
	if errValidation != nil {
		return response.CreatedSocialMediaResponse{}, errValidation
	}
	sm.Created_At = time.Now().Format("2006-01-02")
	sm.Updated_At = time.Now().Format("2006-01-02")
	// Panggil Repository
	result := s.SocialMediaRepository.Create(sm)
	// Return
	resp := response.CreatedSocialMediaResponse{
		Id:               result.Id,
		Name:             result.Name,
		Social_Media_Url: result.Social_Media_Url,
		User_Id:          result.User_Id,
		Created_At:       result.Created_At,
	}
	return resp, nil
}

// Update implements SocialMediaService.
func (*SocialMediaServiceImpl) Update(sm request.UpdateSocialMediaRequest) (response.UpdatedSocialMediaResponse, error) {
	panic("unimplemented")
}

func NewSocialMediaServiceImpl(sm repository.SocialMediaRepository, v *validator.Validate) SocialMediaService {
	return &SocialMediaServiceImpl{
		SocialMediaRepository: sm,
		Validate:              v,
	}
}
