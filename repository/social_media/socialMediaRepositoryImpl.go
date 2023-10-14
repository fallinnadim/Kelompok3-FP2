package repository

import (
	"database/sql"
	"errors"
	request "fp2/data/request/social_media"
	response "fp2/data/response/social_media"
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
	query := `
		DELETE FROM social_medias WHERE id = $1;
	`
	s.Db.Exec(query, id)
}

// FindAll implements SocialMediaRepository.
func (s *SocialMediaRepositoryImpl) FindAll() []response.AllSocialMediaResponse {
	socialMedia := []response.AllSocialMediaResponse{}
	query := `
		SELECT sm.id, sm.name, sm.social_media_url, sm.user_id, sm.created_at, sm.updated_at, u.id, u.username, u.email
		FROM social_medias AS sm
		JOIN users AS u ON sm.user_id = u.id;
	`
	rows, _ := s.Db.Query(query)
	defer rows.Close()
	for rows.Next() {
		sosMed := response.AllSocialMediaResponse{}
		rows.Scan(&sosMed.Id, &sosMed.Name, &sosMed.Social_Media_Url, &sosMed.User_Id, &sosMed.Created_At, &sosMed.Updated_At, &sosMed.User.Id, &sosMed.User.Username, &sosMed.User.Email)
		socialMedia = append(socialMedia, sosMed)
	}
	return socialMedia
}

// Update implements SocialMediaRepository.
func (s *SocialMediaRepositoryImpl) Update(sm request.UpdateSocialMediaRequest) models.SocialMedia {
	var updatedResult = models.SocialMedia{}
	query := `
		UPDATE social_medias
		SET name = $1, social_media_url = $2, updated_at = $3
		WHERE id = $4
		RETURNING *;
	`
	s.Db.QueryRow(query, sm.Name, sm.Social_Media_Url, sm.Updated_At, sm.Id).Scan(&updatedResult.Id, &updatedResult.Name, &updatedResult.Social_Media_Url, &updatedResult.User_Id, &updatedResult.Created_At, &updatedResult.Updated_At)

	return updatedResult
}

// Update implements SocialMediaRepository.
func (s *SocialMediaRepositoryImpl) FindById(smId int) (models.SocialMedia, error) {
	var socialMedia = models.SocialMedia{}
	query := `
		SELECT * FROM social_medias WHERE id = $1;
	`
	errQuery := s.Db.QueryRow(query, smId).Scan(&socialMedia.Id, &socialMedia.Name, &socialMedia.Social_Media_Url, &socialMedia.User_Id, &socialMedia.Created_At, &socialMedia.Updated_At)
	if errQuery == sql.ErrNoRows {
		return socialMedia, errors.New("Social media not found")
	}
	return socialMedia, nil
}

func NewSocialMediaRepositoryImpl(db *sql.DB) SocialMediaRepository {
	return &SocialMediaRepositoryImpl{Db: db}
}
