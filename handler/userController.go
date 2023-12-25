package handler

import (
	"fp2/dto"
	"fp2/helper"
	services "fp2/services/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ambil Body request
// Return response
type UserController struct {
	UserService services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{UserService: s}
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Invalid token",
		})
		return
	}
	// panggil service
	updateUserRequest := dto.UpdateUserRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
	if err != nil {
		statusCode, errMessage := helper.ParseError(err)
		webResponse := dto.FailedResponse{
			Status:  false,
			Message: errMessage,
		}
		ctx.JSON(statusCode, webResponse)
		return
	}
	updateUserRequest.Id = userId.(int)
	result, errUpdate := u.UserService.Update(updateUserRequest)
	// return response
	if errUpdate != nil {
		statusCode, errMessage := helper.ParseError(errUpdate)
		webResponse := dto.FailedResponse{
			Status:  false,
			Message: errMessage,
		}
		ctx.JSON(statusCode, webResponse)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Invalid token",
		})
		return
	}
	// panggil service
	u.UserService.Delete(userId.(int))
	// return response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
