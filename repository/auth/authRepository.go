package repository

import "fp2/models"

type AuthRepository interface {
	Create(user models.User) error
	FindEmail(email string) (user models.User, err error)
}
