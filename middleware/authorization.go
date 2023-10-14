package middleware

import (
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
