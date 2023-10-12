package services

import (
	"fp2/data/request"
	"fp2/data/response"
)

type UserService interface {
	Update(user request.UpdateUserRequest) (response.UpdatedUserResponse, error)
	Delete(id int)
}
