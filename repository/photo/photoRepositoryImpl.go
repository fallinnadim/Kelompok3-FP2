package repository

import (
	"database/sql"
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
func (*PhotoRepositoryImpl) Delete(id int) {
	panic("unimplemented")
}

// FindAll implements PhotoRepository.
func (*PhotoRepositoryImpl) FindAll() []response.AllPhotoResponse {
	panic("unimplemented")
}

// FindById implements PhotoRepository.
func (*PhotoRepositoryImpl) FindById(id int) (models.Photo, error) {
	panic("unimplemented")
}

// Update implements PhotoRepository.
func (*PhotoRepositoryImpl) Update(sm request.UpdatePhotoRequest) models.Photo {
	panic("unimplemented")
}

func NewPhotoRepositoryImpl(db *sql.DB) PhotoRepository {
	return &PhotoRepositoryImpl{Db: db}
}
