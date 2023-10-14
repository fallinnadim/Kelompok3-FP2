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
func (*PhotoRepositoryImpl) Create(sm request.CreatePhotoRequest) models.SocialMedia {
	panic("unimplemented")
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
func (*PhotoRepositoryImpl) Update(sm request.UpdatePhotoRequest) models.SocialMedia {
	panic("unimplemented")
}

func NewPhotoRepositoryImpl(db *sql.DB) PhotoRepository {
	return &PhotoRepositoryImpl{Db: db}
}
