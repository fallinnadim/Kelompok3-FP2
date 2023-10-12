package services

import (
	"fp2/data/request/users"
	"fp2/data/response/users"
)

type UserService interface {
	Update(user request.UpdateUserRequest) (response.UpdatedUserResponse, error)
	Delete(id int)
}
