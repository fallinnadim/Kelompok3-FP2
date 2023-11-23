package services

import "fp2/dto"

type AuthService interface {
	Register(user dto.CreateUserRequest) (dto.CreatedUserResponse, error)
	CheckEmail(email string) error
	CheckUsername(username string) error
	Login(user dto.LoginUserRequest) (string, error)
}
