package services

import (
	"fp2/data/request"
	"fp2/models"
)

type AuthService interface {
	Register(user request.CreateUserRequest) (models.User, error)
	CheckEmail(email string) error
	CheckUsername(username string) error
	Login(user request.LoginUserRequest) (string, error)
}
