package services

import "fp2/dto"

type UserService interface {
	Update(user dto.UpdateUserRequest) (dto.UpdatedUserResponse, error)
	Delete(id int)
}
