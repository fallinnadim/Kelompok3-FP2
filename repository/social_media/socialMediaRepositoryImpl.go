package repository

import (
	"database/sql"
	request "fp2/data/request/social_media"
	"fp2/models"
)

type SocialMediaRepositoryImpl struct {
	Db *sql.DB
}

// Create implements SocialMediaRepository.
func (s *SocialMediaRepositoryImpl) Create(sm request.CreateSocialMediaRequest) models.SocialMedia {
	var newSocialMedia = models.SocialMedia{}
	query := `
		INSERT INTO social_medias (name, social_media_url, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *;
	`
	s.Db.QueryRow(query, sm.Name, sm.Social_Media_Url, sm.User_Id, sm.Created_At, sm.Updated_At).Scan(&newSocialMedia.Id, &newSocialMedia.Name, &newSocialMedia.Social_Media_Url, &newSocialMedia.User_Id, &newSocialMedia.Created_At, &newSocialMedia.Updated_At)

	return newSocialMedia
}

// Delete implements SocialMediaRepository.
func (s *SocialMediaRepositoryImpl) Delete(id int) {
	panic("unimplemented")
}

// FindAll implements SocialMediaRepository.
func (s *SocialMediaRepositoryImpl) FindAll(id int) []models.SocialMedia {
	panic("unimplemented")
}

// Update implements SocialMediaRepository.
func (s *SocialMediaRepositoryImpl) Update(sm request.UpdateSocialMediaRequest) models.SocialMedia {
	panic("unimplemented")
}

func NewSocialMediaRepositoryImpl(db *sql.DB) SocialMediaRepository {
	return &SocialMediaRepositoryImpl{Db: db}
}
