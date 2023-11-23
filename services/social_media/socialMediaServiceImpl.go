package services

import (
	"fp2/dto"
	repository "fp2/repository/social_media"
	"time"

	"github.com/go-playground/validator/v10"
)

type SocialMediaServiceImpl struct {
	SocialMediaRepository repository.SocialMediaRepository
	Validate              *validator.Validate
}

// Delete implements SocialMediaService.
func (s *SocialMediaServiceImpl) Delete(smId int) error {
	// Panggil service
	s.SocialMediaRepository.Delete(smId)
	return nil
}

// GetAll implements SocialMediaService.
func (s *SocialMediaServiceImpl) GetAll() []dto.AllSocialMediaResponse {
	result := s.SocialMediaRepository.FindAll()
	return result
}

// Post implements SocialMediaService.
func (s *SocialMediaServiceImpl) Post(sm dto.CreateSocialMediaRequest) (dto.CreatedSocialMediaResponse, error) {
	// Validasi Struct
	errValidation := s.Validate.Struct(sm)
	if errValidation != nil {
		return dto.CreatedSocialMediaResponse{}, errValidation
	}
	sm.Created_At = time.Now().Format("2006-01-02")
	sm.Updated_At = time.Now().Format("2006-01-02")
	// Panggil Repository
	result := s.SocialMediaRepository.Create(sm)
	// Return
	resp := dto.CreatedSocialMediaResponse{
		Id:               result.Id,
		Name:             result.Name,
		Social_Media_Url: result.Social_Media_Url,
		User_Id:          result.User_Id,
		Created_At:       result.Created_At,
	}
	return resp, nil
}

// Update implements SocialMediaService.
func (s *SocialMediaServiceImpl) Update(sm dto.UpdateSocialMediaRequest) (dto.UpdatedSocialMediaResponse, error) {
	// Validasi Struct
	errValidation := s.Validate.Struct(sm)
	if errValidation != nil {
		return dto.UpdatedSocialMediaResponse{}, errValidation
	}
	sm.Updated_At = time.Now().Format("2006-01-02")
	// Panggil service
	result := s.SocialMediaRepository.Update(sm)
	updateSocialMedia := dto.UpdatedSocialMediaResponse{
		Id:               result.Id,
		Name:             result.Name,
		Social_Media_Url: result.Social_Media_Url,
		User_Id:          result.User_Id,
		Updated_At:       result.Updated_At,
	}
	return updateSocialMedia, nil
}

func NewSocialMediaServiceImpl(sm repository.SocialMediaRepository, v *validator.Validate) SocialMediaService {
	return &SocialMediaServiceImpl{
		SocialMediaRepository: sm,
		Validate:              v,
	}
}
