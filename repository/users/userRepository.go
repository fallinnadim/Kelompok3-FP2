package repository

import "fp2/models"

type UserRepository interface {
	FindById(id int) (user models.User, err error)
	Update(user models.User)
	Delete(id int)
}
