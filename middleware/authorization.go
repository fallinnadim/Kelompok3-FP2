package middleware

import (
	"fp2/helper"
	cRepository "fp2/repository/comment"
	pRepository "fp2/repository/photo"
	repository "fp2/repository/social_media"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AuthorizedUserSm(smRepository repository.SocialMediaRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, _ := ctx.Get("userId")
		smId, _ := strconv.Atoi(ctx.Param("socialMediaId"))
		result, errFind := smRepository.FindById(smId)
		if errFind != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Social Media Not found",
			})
			return
		}
		if result.User_Id != userId {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		ctx.Next()
	}
}

func AuthorizedUserP(p pRepository.PhotoRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, _ := ctx.Get("userId")
		photoId, _ := strconv.Atoi(ctx.Param("photoId"))
		result, errFind := p.FindById(photoId)
		if errFind != nil {
			statusCode, errMessage := helper.ParseError(errFind)
			ctx.AbortWithStatusJSON(statusCode, gin.H{
				"message": errMessage,
			})
			return
		}
		if result.User_Id != userId {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		ctx.Next()
	}
}

func AuthorizedUserC(c cRepository.CommentRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, _ := ctx.Get("userId")
		commentId, _ := strconv.Atoi(ctx.Param("commentId"))
		result, errFind := c.FindById(commentId)
		if errFind != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Comment Not found",
			})
			return
		}
		if result.User_Id != userId {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		ctx.Next()
	}
}
