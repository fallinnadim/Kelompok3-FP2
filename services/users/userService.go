package services

import (
	"fp2/data/request"
	"fp2/data/response"
)

type UserService interface {
	Update(user request.UpdateUserRequest) error
	Delete(id int) error
	FindById(user request.LoginUserRequest) (response.UserSuccessRespons, error)
}
