package controller

import (
	"fp2/data/request/users"
	"fp2/data/response/users"
	"fp2/helper"
	services "fp2/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	AuthenticationService services.AuthService
}

func NewAuthenticationController(s services.AuthService) *AuthenticationController {
	return &AuthenticationController{AuthenticationService: s}
}

func (a *AuthenticationController) Login(ctx *gin.Context) {
	loginRequest := request.LoginUserRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		statusCode, errMessage := helper.ParseError(err)
		webResponse := response.FailedResponse{
			Status:  false,
			Message: errMessage,
		}
		ctx.JSON(statusCode, webResponse)
		return
	}
	token, errToken := a.AuthenticationService.Login(loginRequest)
	if errToken != nil {
		statusCode, errMessage := helper.ParseError(errToken)
		webResponse := response.FailedResponse{
			Status:  false,
			Message: errMessage,
		}
		ctx.JSON(statusCode, webResponse)
		return
	}
	webResponse := response.LoginResponse{
		Token: token,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (a *AuthenticationController) Register(ctx *gin.Context) {
	createUserRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	if err != nil {
		statusCode, errMessage := helper.ParseError(err)
		webResponse := response.FailedResponse{
			Status:  false,
			Message: errMessage,
		}
		ctx.JSON(statusCode, webResponse)
		return
	}
	newUser, errRegister := a.AuthenticationService.Register(createUserRequest)
	if errRegister != nil {
		statusCode, errMessage := helper.ParseError(errRegister)
		webResponse := response.FailedResponse{
			Status:  false,
			Message: errMessage,
		}
		ctx.JSON(statusCode, webResponse)
		return
	}
	ctx.JSON(http.StatusCreated, newUser)
}
