package services

import (
	"fp2/data/request"
	"fp2/data/response"
	repository "fp2/repository/users"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	validate       *validator.Validate
}

// Delete implements UserService.
func (*UserServiceImpl) Delete(id int) error {
	panic("unimplemented")
}

// FindById implements UserService.
func (*UserServiceImpl) FindById(user request.LoginUserRequest) (response.UserSuccessRespons, error) {
	panic("unimplemented")
}

// Update implements UserService.
func (*UserServiceImpl) Update(user request.UpdateUserRequest) error {
	panic("unimplemented")
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		validate:       validate,
	}
}
