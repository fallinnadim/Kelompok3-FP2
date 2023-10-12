package controller

import (
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

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Update",
	})
}

func (u *UserController) DeleteUser(ctx *gin.Context) {

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Delete",
	})
}
