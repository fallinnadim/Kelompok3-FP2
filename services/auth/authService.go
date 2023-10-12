package services

import (
	"fp2/data/request/users"
	"fp2/data/response/users"
)

type AuthService interface {
	Register(user request.CreateUserRequest) (response.CreatedUserResponse, error)
	CheckEmail(email string) error
	CheckUsername(username string) error
	Login(user request.LoginUserRequest) (string, error)
}
