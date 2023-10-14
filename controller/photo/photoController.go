package controller

import (
	"fmt"
	services "fp2/services/photo"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	PhotoService services.PhotoService
}

func NewPhotoController(s services.PhotoService) *PhotoController {
	return &PhotoController{PhotoService: s}
}

func (p *PhotoController) CreatePhoto(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Invalid token",
		})
		return
	}
	fmt.Println(userId.(int))
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Created",
	})
}

func (p *PhotoController) GetAllPhoto(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get all",
	})
}

func (p *PhotoController) UpdatePhoto(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update",
	})
}

func (p *PhotoController) DeletePhoto(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete",
	})
}
