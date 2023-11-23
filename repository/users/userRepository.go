package repository

import (
	"fp2/dto"
	"fp2/entity"
)

type UserRepository interface {
	FindById(id int) (user entity.User, err error)
	Update(user dto.UpdateUserRequest) entity.User
	Delete(id int)
}
