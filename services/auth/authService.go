package services

import (
	"fp2/data/request"
)

type AuthService interface {
	Register(user request.CreateUserRequest) error
	Login(user request.LoginUserRequest) (string, error)
}
