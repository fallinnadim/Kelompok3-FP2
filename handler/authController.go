package handler

import (
	"fp2/dto"
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
	loginRequest := dto.LoginUserRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		statusCode, errMessage := helper.ParseError(err)
		webResponse := dto.FailedResponse{
			Status:  false,
			Message: errMessage,
		}
		ctx.JSON(statusCode, webResponse)
		return
	}
	token, errToken := a.AuthenticationService.Login(loginRequest)
	if errToken != nil {
		statusCode, errMessage := helper.ParseError(errToken)
		webResponse := dto.FailedResponse{
			Status:  false,
			Message: errMessage,
		}
		ctx.JSON(statusCode, webResponse)
		return
	}
	webResponse := dto.LoginResponse{
		Token: token,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func (a *AuthenticationController) Register(ctx *gin.Context) {
	createUserRequest := dto.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	if err != nil {
		statusCode, errMessage := helper.ParseError(err)
		webResponse := dto.FailedResponse{
			Status:  false,
			Message: errMessage,
		}
		ctx.JSON(statusCode, webResponse)
		return
	}
	newUser, errRegister := a.AuthenticationService.Register(createUserRequest)
	if errRegister != nil {
		statusCode, errMessage := helper.ParseError(errRegister)
		webResponse := dto.FailedResponse{
			Status:  false,
			Message: errMessage,
		}
		ctx.JSON(statusCode, webResponse)
		return
	}
	ctx.JSON(http.StatusCreated, newUser)
}
