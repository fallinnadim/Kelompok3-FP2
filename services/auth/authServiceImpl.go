package services

import (
	"errors"
	"fp2/config"
	"fp2/data/request"
	"fp2/data/response"
	"fp2/models"
	repository "fp2/repository/auth"
	"fp2/utils"
	"log"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	validate       *validator.Validate
}

// Login implements AuthService.
func (a *AuthServiceImpl) Login(user request.LoginUserRequest) (string, error) {
	// Validasi Struct
	errValidation := a.validate.Struct(user)
	if errValidation != nil {
		return "", errValidation
	}
	loginUser, errUser := a.AuthRepository.FindEmail(user.Email)
	if errUser != nil {
		return "", errors.New("Invalid username or password")
	}
	config.LoadConfig()
	verifyError := utils.VerifyPassword(loginUser.Password, user.Password)
	if verifyError != nil {
		return "", errors.New("Invalid username or password")
	}
	// Generate Token
	token, errToken := utils.GenerateToken(time.Minute*60, loginUser.Id, os.Getenv("TOKEN_SECRET"))
	if errToken != nil {
		log.Fatalln(errToken)
	}
	return token, nil
}

// Register implements AuthService.
func (a *AuthServiceImpl) Register(user request.CreateUserRequest) (response.CreatedUserResponse, error) {
	// Validasi Struct
	errValidation := a.validate.Struct(user)
	if errValidation != nil {
		return response.CreatedUserResponse{}, errValidation
	}
	// Cek Email
	if err := a.CheckEmail(user.Email); err == nil { // err nil -> artinya email ketemu, return disini
		return response.CreatedUserResponse{}, errors.New("Silahkan gunakan Email lain")
	}
	// Cek Username
	if err := a.CheckUsername(user.Username); err == nil { // err nil -> artinya username ketemu, return disini
		return response.CreatedUserResponse{}, errors.New("Silahkan gunakan Username lain")
	}
	// Lewat dari sini email dan username available
	hashedPassword, _ := utils.HashPassword(user.Password)
	newUser := models.User{
		Username:   user.Username,
		Email:      user.Email,
		Password:   hashedPassword,
		Age:        user.Age,
		Created_At: time.Now().Format("2006-01-02"),
		Updated_At: time.Now().Format("2006-01-02"),
	}
	result := a.AuthRepository.Create(newUser)
	createdUser := response.CreatedUserResponse{
		Id:       result.Id,
		Email:    result.Email,
		Username: result.Username,
		Age:      result.Age,
	}
	return createdUser, nil
}

// Check EMail AuthService.
func (a *AuthServiceImpl) CheckEmail(email string) error {
	_, err := a.AuthRepository.FindEmail(email)
	// kalau error artinya tidak ketemu dan email available
	return err
}

// Check Username AuthService.
func (a *AuthServiceImpl) CheckUsername(username string) error {
	_, err := a.AuthRepository.FindUsername(username)
	// kalau error artinya tidak ketemu dan username available
	return err
}

func NewAuthServiceImpl(a repository.AuthRepository, v *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: a,
		validate:       v,
	}
}
