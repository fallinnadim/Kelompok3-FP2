package services

import (
	"errors"
	"fp2/data/request"
	"fp2/data/response"
	authRepository "fp2/repository/auth"
	repository "fp2/repository/users"
	"time"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	AuthRepository authRepository.AuthRepository
	validate       *validator.Validate
}

// Delete implements UserService.
func (u *UserServiceImpl) Delete(id int) {
	u.UserRepository.Delete(id)
}

// Update implements UserService.
func (u *UserServiceImpl) Update(user request.UpdateUserRequest) (response.UpdatedUserResponse, error) {
	// Validasi Struct
	errValidation := u.validate.Struct(user)
	if errValidation != nil {
		return response.UpdatedUserResponse{}, errValidation
	}
	// find by id
	resultId, _ := u.UserRepository.FindById(user.Id)
	// Cek email
	_, errEmail := u.AuthRepository.FindEmail(user.Email)
	if errEmail == nil && user.Email != resultId.Email { // artinya email sudah dipakai pada record lain
		return response.UpdatedUserResponse{}, errors.New("Email tidak bisa dipakai")
	}
	// Cek username
	_, errUsername := u.AuthRepository.FindUsername(user.Username)
	if errUsername == nil && user.Username != resultId.Username { // artinya email sudah dipakai
		return response.UpdatedUserResponse{}, errors.New("Username tidak bisa dipakai")
	}
	user.Updated_At = time.Now().Format("2006-01-02")
	result := u.UserRepository.Update(user)
	updatedUser := response.UpdatedUserResponse{
		Id:         result.Id,
		Email:      result.Email,
		Username:   result.Username,
		Age:        result.Age,
		Updated_At: result.Updated_At,
	}
	return updatedUser, nil
}

func NewUserServiceImpl(userRepository repository.UserRepository, authRepository authRepository.AuthRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		AuthRepository: authRepository,
		validate:       validate,
	}
}
