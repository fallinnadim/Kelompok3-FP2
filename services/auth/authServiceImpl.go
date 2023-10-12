package services

import (
	"errors"
	"fp2/config"
	"fp2/data/request"
	"fp2/helper"
	"fp2/models"
	repository "fp2/repository/auth"
	"fp2/utils"
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
	// validasi struct
	newUser, errUser := a.AuthRepository.FindEmail(user.Email)
	if errUser != nil {
		return "", errors.New("Invalid username or password")
	}
	config.LoadConfig()
	verifyError := utils.VerifyPassword(newUser.Password, user.Password)
	if verifyError != nil {
		return "", errors.New("Invalid username or password")
	}
	// Generate Token
	token, errToken := utils.GenerateToken(time.Minute*60, newUser.Id, os.Getenv("TOKEN_SECRET"))
	helper.ErrorFatal(errToken)
	return token, nil
}

// Register implements AuthService.
func (a *AuthServiceImpl) Register(user request.CreateUserRequest) error {
	// validasi struct
	hashedPassword, err := utils.HashPassword(user.Password)
	helper.ErrorFatal(err)
	newUser := models.User{
		Username:   user.Username,
		Email:      user.Email,
		Password:   hashedPassword,
		Age:        user.Age,
		Created_At: time.Now().Format("2006-01-02"),
		Updated_At: time.Now().Format("2006-01-02"),
	}
	a.AuthRepository.Create(newUser)
	return nil
}

func NewAuthServiceImpl(a repository.AuthRepository, v *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: a,
		validate:       v,
	}
}
