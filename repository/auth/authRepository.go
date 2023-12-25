package repository

import "fp2/entity"

type AuthRepository interface {
	Create(user entity.User) entity.User
	FindEmail(email string) (user entity.User, err error)
	FindUsername(username string) (user entity.User, err error)
}
