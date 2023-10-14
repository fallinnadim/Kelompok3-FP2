package repository

import (
	"database/sql"
	"errors"
	request "fp2/data/request/photo"
	response "fp2/data/response/photo"
	"fp2/models"
)

type PhotoRepositoryImpl struct {
	Db *sql.DB
}

// Create implements PhotoRepository.
func (p *PhotoRepositoryImpl) Create(cp request.CreatePhotoRequest) models.Photo {
	var newPhoto = models.Photo{}
	query := `
		INSERT INTO photos (title, caption, photo_url, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *;
	`
	p.Db.QueryRow(query, cp.Title, cp.Caption, cp.Photo_Url, cp.User_Id, cp.Created_At, cp.Updated_At).Scan(&newPhoto.Id, &newPhoto.Title, &newPhoto.Caption, &newPhoto.Photo_Url, &newPhoto.User_Id, &newPhoto.Created_At, &newPhoto.Updated_At)

	return newPhoto
}

// Delete implements PhotoRepository.
func (p *PhotoRepositoryImpl) Delete(id int) {
	query := `
		DELETE FROM photos WHERE id = $1;
	`
	p.Db.Exec(query, id)
}

// FindAll implements PhotoRepository.
func (p *PhotoRepositoryImpl) FindAll() []response.AllPhotoResponse {
	photos := []response.AllPhotoResponse{}
	query := `
		SELECT p.id, p.title, p.caption, p.photo_url, p.user_id, p.created_at, p.updated_at, u.id, u.username, u.email
		FROM photos AS p
		JOIN users AS u ON p.user_id = u.id;
	`
	rows, _ := p.Db.Query(query)
	defer rows.Close()
	for rows.Next() {
		photo := response.AllPhotoResponse{}
		rows.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.Photo_Url, &photo.User_Id, &photo.Created_At, &photo.Updated_At, &photo.User.Id, &photo.User.Username, &photo.User.Email)
		photos = append(photos, photo)
	}
	return photos
}

// FindById implements PhotoRepository.
func (p *PhotoRepositoryImpl) FindById(id int) (models.Photo, error) {
	var photo = models.Photo{}
	query := `
		SELECT * FROM photos WHERE id = $1;
	`
	errQuery := p.Db.QueryRow(query, id).Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.Photo_Url, &photo.User_Id, &photo.Created_At, &photo.Updated_At)
	if errQuery == sql.ErrNoRows {
		return photo, errors.New("Photo not found")
	}
	return photo, nil
}

// Update implements PhotoRepository.
func (p *PhotoRepositoryImpl) Update(up request.UpdatePhotoRequest) models.Photo {
	var updatedResult = models.Photo{}
	query := `
		UPDATE photos
		SET title = $1, caption = $2, photo_url = $3, updated_at = $4
		WHERE id = $5
		RETURNING *;
	`
	p.Db.QueryRow(query, up.Title, up.Caption, up.Photo_Url, up.Updated_At, up.Id).Scan(&updatedResult.Id, &updatedResult.Title, &updatedResult.Caption, &updatedResult.Photo_Url, &updatedResult.User_Id, &updatedResult.Created_At, &updatedResult.Updated_At)

	return updatedResult
}

func NewPhotoRepositoryImpl(db *sql.DB) PhotoRepository {
	return &PhotoRepositoryImpl{Db: db}
}
