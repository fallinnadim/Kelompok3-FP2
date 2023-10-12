package repository

import "fp2/models"

type AuthRepository interface {
	Create(user models.User) models.User
	FindEmail(email string) (user models.User, err error)
	FindUsername(username string) (user models.User, err error)
}
