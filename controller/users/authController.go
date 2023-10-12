package controller

import (
	"fp2/data/request"
	"fp2/data/response"
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
		webResponse := response.FailedResponse{
			Status:  false,
			Message: helper.ParseError(err),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	token, errToken := a.AuthenticationService.Login(loginRequest)
	if errToken != nil {
		webResponse := response.FailedResponse{
			Status:  false,
			Message: helper.ParseError(errToken),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
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
		webResponse := response.FailedResponse{
			Status:  false,
			Message: helper.ParseError(err),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	newUser, errRegister := a.AuthenticationService.Register(createUserRequest)
	if errRegister != nil {
		webResponse := response.FailedResponse{
			Status:  false,
			Message: helper.ParseError(errRegister),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	webResponse := response.CreatedUserResponse{
		Id:       newUser.Id,
		Email:    newUser.Email,
		Username: newUser.Username,
		Age:      newUser.Age,
	}
	ctx.JSON(http.StatusCreated, webResponse)
}
