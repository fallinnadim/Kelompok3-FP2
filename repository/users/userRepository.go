package repository

import (
	"fp2/data/request"
	"fp2/models"
)

type UserRepository interface {
	FindById(id int) (user models.User, err error)
	Update(user request.UpdateUserRequest) models.User
	Delete(id int)
}
