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
	helper.ErrorFatal(err)

	token, errToken := a.AuthenticationService.Login(loginRequest)
	if errToken != nil {
		webResponse := response.Response{
			Status:  false,
			Message: "Invalid Username or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}
	webResponse := response.Response{
		Status:  true,
		Message: "Successfully Login",
		Data:    resp,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (a *AuthenticationController) Register(ctx *gin.Context) {
	createUserRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helper.ErrorFatal(err)
	a.AuthenticationService.Register(createUserRequest)
	webResponse := response.Response{
		Status:  true,
		Message: "Successfully Created User",
		Data:    nil,
	}
	ctx.JSON(http.StatusCreated, webResponse)
}
